package config

import configx "github.com/iamKienb/go-core/config"

type ShopQueryConfig struct {
	ES     configx.ElasticSearchConfig `envPrefix:"SHOP_QUERY_SERVICE"`
	Server configx.Server              `envPrefix:"SHOP_QUERY_SERVICE"`
}
