package shop

import (
	"user-command-module/internal/application/services/shop/i18n"
	"user-command-module/internal/domain/member"
	"user-command-module/internal/domain/shop"

	"github.com/iamKienb/go-core/app_error"
)

var shopErrorMap = app_error.ServiceErrorMap{
	shop.ErrShopSlugTaken:      {Kind: app_error.KindValidation, Msg: i18n.MsgSlugTaken},
	shop.ErrShopConflict:       {Kind: app_error.KindConflict, Msg: i18n.MsgCreateShopToFast},
	shop.ErrShopNotFound:       {Kind: app_error.KindNotFound, Msg: i18n.MsgShopNotFound},
	shop.ErrShopInvalid:        {Kind: app_error.KindValidation, Msg: i18n.MsgShopInvalid},
	shop.ErrAddressTypeInvalid: {Kind: app_error.KindValidation, Msg: i18n.MsgAddressTypeInvalid},

	member.ErrActionNotDefined: {Kind: app_error.KindValidation, Msg: i18n.MsgActionInvalid},
	member.ErrShopDenied:       {Kind: app_error.KindForbidden, Msg: i18n.MsgShopDenied},
	member.ErrProductDenied:    {Kind: app_error.KindForbidden, Msg: i18n.MsgProductDenied},
	member.ErrInventoryDenied:  {Kind: app_error.KindForbidden, Msg: i18n.MsgInventoryDenied},
}

func (s *shopService) wrapError(err error) error {
	return app_error.WrapError(err, shopErrorMap)
}
