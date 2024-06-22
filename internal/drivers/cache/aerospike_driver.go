package cache

import (
	"fmt"
	as "github.com/aerospike/aerospike-client-go/v7"
	rrdServerGen "github.com/lissdx/aerospike-ha/internal/pkg/gen/openapi/rrd_server_gen"
	"github.com/lissdx/aerospike-ha/internal/pkg/logger"
	"github.com/lissdx/aerospike-ha/internal/pkg/process"
	"github.com/spf13/viper"
)

var _ process.Processor = (*AerospikeDriver)(nil)

const Namespace = "test"
const Set = "metrics"

type MetricType string

const defaultMaxRecordsFetch = 1000

const (
	NetworkMetricType     MetricType = "network"
	BandwidthMetricType   MetricType = "bandwidth"
	TemperatureMetricType MetricType = "temperature"
	CPULoadMetricType     MetricType = "cpu"
)

func (mt MetricType) IsValid() bool {
	switch mt {
	case NetworkMetricType, BandwidthMetricType, TemperatureMetricType, CPULoadMetricType:
		return true
	default:
		return false
	}
}

type AerospikeDriver struct {
	client *as.Client
	logger logger.ILogger
	host   string
	port   int
}

func (asd *AerospikeDriver) Run() {
	client, err := as.NewClient(asd.host, asd.port)

	if err != nil {
		asd.logger.Panic(err)
	}

	asd.client = client
}

func (asd *AerospikeDriver) Stop() {
	asd.client.Close()
}

func (asd *AerospikeDriver) PutMetric(metricsData rrdServerGen.MetricsData) error {
	if !MetricType(metricsData.Type).IsValid() {
		return fmt.Errorf("PutMetric invalid metric type: %s", metricsData.Type)
	}
	// Create key for a record
	key, err := as.NewKey(Namespace, Set, fmt.Sprintf("%d_%s", int64(metricsData.Timestamp), metricsData.Type))
	if err != nil {
		return fmt.Errorf("cannot create aerospike key for metricsData %+v: error: %w for ", metricsData, err)
	}

	binMetric := as.NewBin("metric", metricsData.Metric)
	binTimestamp := as.NewBin("timestamp", metricsData.Timestamp)
	binType := as.NewBin("type", metricsData.Type)
	err = asd.client.PutBins(nil, key, binTimestamp, binMetric, binType)

	if err != nil {
		return fmt.Errorf("on PutMetric error: %w", err)
	}

	return nil
}

func (asd *AerospikeDriver) GetMetrics(metricType *string, startTime *float64, endTime *float64, maxFetch *int64) ([]rrdServerGen.MetricsData, error) {
	if metricType != nil && !MetricType(*metricType).IsValid() {
		return nil, fmt.Errorf("GetMetrics invalid metric type: %s", *metricType)
	}

	if startTime != nil && *startTime < 0 {
		return nil, fmt.Errorf("GetMetrics invalid start time: %d", int64(*startTime))
	}

	if endTime != nil && *endTime < 0 {
		return nil, fmt.Errorf("GetMetrics invalid start time: %d", int64(*endTime))
	}

	if startTime != nil && endTime != nil && *startTime > *endTime {
		return nil, fmt.Errorf("GetMetrics invalid startTime -> endTime time-range: %d -> %d", int64(*startTime), int64(*endTime))
	}

	return asd.getGetMetricsHelper(metricType, startTime, endTime, maxFetch)
}

func (asd *AerospikeDriver) getGetMetricsHelper(metricType *string, startTime *float64, endTime *float64, maxFetch *int64) ([]rrdServerGen.MetricsData, error) {
	stmt := as.NewStatement(Namespace, Set)
	policy := as.NewQueryPolicy()

	// init MaxRecords
	policy.MaxRecords = func() int64 {
		if maxFetch != nil && *maxFetch > 0 {
			return *maxFetch
		}
		return defaultMaxRecordsFetch
	}()

	// expression list
	fExpressions := make([]*as.Expression, 0, 3)

	if metricType != nil {
		fExpressions = append(fExpressions, as.ExpEq(as.ExpStringBin("type"), as.ExpStringVal(*metricType)))
	}

	if startTime != nil {
		fExpressions = append(fExpressions, as.ExpGreaterEq(as.ExpFloatBin("timestamp"), as.ExpFloatVal(*startTime)))
	}

	if endTime != nil {
		fExpressions = append(fExpressions, as.ExpLess(as.ExpFloatBin("timestamp"), as.ExpFloatVal(*endTime)))
	}

	if len(fExpressions) == 1 {
		policy.FilterExpression = fExpressions[0]
	} else if len(fExpressions) > 1 {
		policy.FilterExpression = as.ExpAnd(fExpressions...)
	}

	recordset, err := asd.client.Query(policy, stmt)

	if err != nil {
		return nil, fmt.Errorf("on GetMetrics error: %w", err)
	}

	result := make([]rrdServerGen.MetricsData, 0, policy.MaxRecords)
	for rec := range recordset.Results() {
		if rec.Err != nil {
			// if there was an error, handle it if needed
			// Scans are retried in Aerospike servers v5+
			asd.logger.Error(fmt.Errorf("on recordset.Results error: %w", rec.Err))
			continue
		}

		metricsData := recordBinsToMetricsData(rec.Record.Bins)

		result = append(result, *metricsData)
	}

	return result, nil
}

/**
 * Helpers part
 */
func recordBinsToMetricsData(bMap as.BinMap) *rrdServerGen.MetricsData {
	return &rrdServerGen.MetricsData{
		Type:      rrdServerGen.MetricsDataType(bMap["type"].(string)),
		Metric:    bMap["metric"].(float64),
		Timestamp: bMap["timestamp"].(float64)}
}

/**
 * Constructor part
 */
func NewAerospikeDriver(config *viper.Viper, logger logger.ILogger) *AerospikeDriver {

	return &AerospikeDriver{
		client: nil,
		logger: logger,
		host:   config.GetString("AS_DB_DRIVER_HOST"),
		port:   config.GetInt("AS_DB_DRIVER_PORT"),
	}
}
