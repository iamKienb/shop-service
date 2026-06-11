package module

import (
	"shop-worker-module/internal/application/port"
	"shop-worker-module/internal/application/processor"
)

type ApplicationModule struct {
	EventProcessor port.EventProcessor
}

func NewApplicationModule(infra *InfraModule) *ApplicationModule {
	return &ApplicationModule{
		EventProcessor: processor.NewShopEventProcessor(infra.ESRepo, infra.workerCache),
	}
}
