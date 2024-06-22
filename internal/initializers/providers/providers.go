package providers

func Dependencies() []interface{} {
	return []interface{}{
		NewAppConfig,
		NewProcessConfig,
		NewLogger,
		NewAerospikeDriver,
		NewRrdService,
		ProcessorFactory,
	}
}
