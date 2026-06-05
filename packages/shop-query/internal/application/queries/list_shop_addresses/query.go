package list_shop_addresses

import (
	"context"

	"shop-query-module/internal/application/service/models"
)

type Query struct{ ShopID string }
type Result struct{ Addresses []models.ShopAddress }
type Executor interface {
	Execute(context.Context, Query) (*Result, error)
}
