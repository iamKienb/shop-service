package member

import "errors"

var (
	ErrShopInvalid        = errors.New("shop_invalid")
	ErrAddressTypeInvalid = errors.New("shop_address_type_invalid")
	ErrShopSlugTaken      = errors.New("shop_slug_already_taken")
	ErrShopConflict       = errors.New("create_shop_too_fast")
	ErrShopNotFound       = errors.New("shop_not_found")

	ErrActionNotDefined = errors.New("action_not_defined")
	ErrShopDenied       = errors.New("shop_permission_denied")
	ErrProductDenied    = errors.New("product_permission_denied")
	ErrInventoryDenied  = errors.New("inventory_permission_denied")
)
