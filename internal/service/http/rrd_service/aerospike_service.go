package rrd_service

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lissdx/aerospike-ha/internal/drivers/cache"
	rrdServerGen "github.com/lissdx/aerospike-ha/internal/pkg/gen/openapi/rrd_server_gen"
	"github.com/lissdx/aerospike-ha/internal/pkg/logger"
	"github.com/lissdx/aerospike-ha/internal/pkg/process"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var _ process.Processor = (*RrdServer)(nil)
var _ rrdServerGen.ServerInterface = (*RrdServer)(nil)

type responseCode int

const (
	responseCodeOK         responseCode = 777200
	responseCodeNotFound   responseCode = 777204
	responseCodeBadRequest responseCode = 777400
)

type RrdServer struct {
	logger     logger.ILogger
	AsDriver   *cache.AerospikeDriver
	echoServer *echo.Echo
	host       string
	port       int
}

func (r *RrdServer) GetMetrics(ctx echo.Context, params rrdServerGen.GetMetricsParams) error {

	getRes, err := r.AsDriver.GetMetrics((*string)(params.Type), params.Start, params.End, params.MaxFetch)
	if err != nil {
		r.logger.Error(fmt.Errorf("cannot process requested fetch metrics error %w", err))
		res := rrdServerGen.BadRequest{
			Msg:    "cannot fetch data from db",
			Status: int(responseCodeBadRequest),
			Detail: err.Error(),
		}
		return ctx.JSON(http.StatusBadRequest, res)
	}

	if len(getRes) <= 0 {
		return ctx.JSON(http.StatusNoContent, getRes)
	}

	return ctx.JSON(http.StatusOK, getRes)

}

func (r *RrdServer) PutMetrics(ctx echo.Context) error {

	// fetch request body
	var metricsData rrdServerGen.MetricsData
	err := ctx.Bind(&metricsData)

	if err != nil {
		return fmt.Errorf("cannot bind metrics data: %w", err)
	}

	// metricsData.Timestamp should be greater than 0
	if metricsData.Timestamp <= 0 {
		bResp := rrdServerGen.BadRequest{
			Status: int(responseCodeBadRequest),
			Detail: fmt.Sprintf("metrics \\'timestamp\\' should be greater than 0. provided metricsData: %+v", metricsData),
			Msg:    "cannot process requested",
		}
		r.logger.Error(fmt.Sprintf("cannot process requested for metricsData: %+v", metricsData))
		return ctx.JSON(http.StatusBadRequest, bResp)
	}

	// store the given metric
	err = r.AsDriver.PutMetric(metricsData)

	// TODO replace http.StatusBadRequest with http.StatusInternalServerError
	if err != nil {
		bResp := rrdServerGen.BadRequest{
			Status: int(responseCodeBadRequest),
			Detail: "on insert DB error",
			Msg:    "cannot process requested",
		}
		r.logger.Error(fmt.Errorf("cannot process PutMetric for metricsData: %+v, error: %w", metricsData, err).Error())
		return ctx.JSON(http.StatusBadRequest, bResp)
	}

	resp := rrdServerGen.OkResponse{
		Msg:    fmt.Sprintf("metrics data object processed: %v", metricsData),
		Status: int(responseCodeOK),
	}

	r.logger.Debug(fmt.Sprintf("metrics data object processed: %v", metricsData))

	return ctx.JSON(http.StatusOK, resp)
}

func (r *RrdServer) Run() {

	r.AsDriver.Run()

	err := r.echoServer.Start(fmt.Sprintf("%s:%d", r.host, r.port))
	if err != nil {
		r.logger.Fatal(err.Error())
	}
}

func (r *RrdServer) Stop() {

	r.AsDriver.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// TODO: check how we can avoid FATAL msg on shutdown
	// {"level":"FATAL","time":"2024-06-22T08:41:49.055+0300","message":"http: Server closed"}
	err := r.echoServer.Shutdown(ctx)
	if err != nil {
		r.logger.Error(err)
	}
}

func NewRrdServer(config *viper.Viper, logger logger.ILogger, driver *cache.AerospikeDriver) *RrdServer {
	res := &RrdServer{
		logger:     logger,
		AsDriver:   driver,
		port:       config.GetInt("RRD_SERVER_PORT"),
		host:       config.GetString("RRD_SERVER_HOST"),
		echoServer: echo.New(),
	}

	// Middleware
	res.echoServer.Use(middleware.Logger())
	res.echoServer.Use(middleware.Recover())
	rrdServerGen.RegisterHandlers(res.echoServer, res)

	return res
}
