package list_shop_members

import (
	"context"

	"shop-query-module/internal/application/port"
)

type Query struct{ ShopID string }
type Result struct{ Members []port.ShopMember }
type Executor interface {
	Execute(context.Context, Query) (*Result, error)
}
