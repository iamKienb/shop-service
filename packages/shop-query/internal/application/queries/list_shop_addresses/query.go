package list_shop_addresses

import (
	"context"

	"shop-query-module/internal/application/port"
)

type Query struct{ ShopID string }
type Result struct{ Addresses []port.ShopAddress }
type Executor interface {
	Execute(context.Context, Query) (*Result, error)
}
