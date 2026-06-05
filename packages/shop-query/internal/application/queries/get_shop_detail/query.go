package get_shop_detail

import (
	"context"

	"shop-query-module/internal/application/port"
)

type Query struct{ ShopID string }
type Result struct{ Shop *port.Shop }
type Executor interface {
	Execute(context.Context, Query) (*Result, error)
}
