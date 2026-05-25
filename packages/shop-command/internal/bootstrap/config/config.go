package config

import (
	configx "github.com/iamKienb/go-core/config"
)

type ShopCommandConfig struct {
	Postgres configx.PostgresConfig `envPrefix:"SHOP_COMMAND_SERVICE"`
	Redis    configx.RedisConfig    `envPrefix:"SHOP_COMMAND_SERVICE"`
	Server   configx.Server         `envPrefix:"SHOP_COMMAND_SERVICE"`
}
