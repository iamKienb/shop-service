package module

import (
	"context"
	"fmt"
	"shop-worker-module/internal/application/port"
	"shop-worker-module/internal/bootstrap/config"
	"shop-worker-module/internal/infra/cache"
	"shop-worker-module/internal/infra/elasticsearch"

	esx "github.com/iamKienb/go-core/elasticsearch"
	kafkax "github.com/iamKienb/go-core/kafka"
	redisx "github.com/iamKienb/go-core/redis"
)

type InfraModule struct {
	ESService    esx.ESXService
	RedisService redisx.RedisXService

	ESRepo      port.ESRepository
	workerCache port.WorkerCache

	Kafka kafkax.KafkaXService
}

func NewInfraModule(ctx context.Context, cfg *config.ShopWorkerConfig) (*InfraModule, error) {
	esService, err := esx.New(cfg.ES)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch: %w", err)
	}
	redisService, err := redisx.New(cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch: %w", err)
	}

	kafka, err := kafkax.New(cfg.Kafka)
	if err != nil {
		return nil, fmt.Errorf("kafka: %w", err)
	}

	return &InfraModule{
		ESService:    esService,
		RedisService: redisService,

		ESRepo:      elasticsearch.NewESRepository(esService, esService.GetClient()),
		workerCache: cache.NewWorkerCache(redisService.GetClient()),
		Kafka:       kafka,
	}, nil
}
