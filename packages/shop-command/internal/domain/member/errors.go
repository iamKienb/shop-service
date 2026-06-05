package member

import "errors"

var (
	ErrActionNotDefined = errors.New("action_not_defined")
	ErrShopDenied       = errors.New("shop_permission_denied")
	ErrProductDenied    = errors.New("product_permission_denied")
	ErrInventoryDenied  = errors.New("inventory_permission_denied")
)
