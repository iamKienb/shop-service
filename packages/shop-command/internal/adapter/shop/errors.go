package shop

import (
	"shop-command-module/internal/application/services/shop/i18n"
	"shop-command-module/internal/domain/member"
	domain_shop "shop-command-module/internal/domain/shop"

	"github.com/iamKienb/go-core/app_error"
)

var shopErrorMap = app_error.ServiceErrorMap{
	domain_shop.ErrShopSlugTaken:      {Kind: app_error.KindValidation, Msg: i18n.MsgSlugTaken},
	domain_shop.ErrShopConflict:       {Kind: app_error.KindConflict, Msg: i18n.MsgCreateShopToFast},
	domain_shop.ErrShopNotFound:       {Kind: app_error.KindNotFound, Msg: i18n.MsgShopNotFound},
	domain_shop.ErrShopInvalid:        {Kind: app_error.KindValidation, Msg: i18n.MsgShopInvalid},
	domain_shop.ErrAddressTypeInvalid: {Kind: app_error.KindValidation, Msg: i18n.MsgAddressTypeInvalid},
	domain_shop.ErrShopNotAllowed:     {Kind: app_error.KindForbidden, Msg: i18n.MsgShopNotAllowed},

	member.ErrActionNotDefined: {Kind: app_error.KindValidation, Msg: i18n.MsgActionInvalid},
	member.ErrShopDenied:       {Kind: app_error.KindForbidden, Msg: i18n.MsgShopDenied},
	member.ErrProductDenied:    {Kind: app_error.KindForbidden, Msg: i18n.MsgProductDenied},
	member.ErrInventoryDenied:  {Kind: app_error.KindForbidden, Msg: i18n.MsgInventoryDenied},
}

func toApplicationError(err error) error {
	return app_error.WrapError(err, shopErrorMap)
}
