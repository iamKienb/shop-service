package config

import (
	configx "github.com/iamKienb/go-core/config"
)

type ShopWorkerConfig struct {
	ES       configx.ElasticSearchConfig `envPrefix:"SHOP_WORKER_SERVICE"`
	Kafka    configx.KafkaConfig         `envPrefix:"SHOP_WORKER_SERVICE"`
	Consumer configx.ConsumerConfig      `envPrefix:"SHOP_WORKER_SERVICE"`
}
