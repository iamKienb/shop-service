package get_shop_detail

import (
	"context"

	"shop-query-module/internal/application/service/models"
)

type Query struct{ ShopID string }
type Result struct{ Shop *models.Shop }
type Executor interface {
	Execute(context.Context, Query) (*Result, error)
}
