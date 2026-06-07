package shop

import "shop-command-module/internal/domain/shared"

type NewShopAddressParams struct {
	UserID shared.UserID
	ShopID shared.ShopID

	CountryID   string
	CountryName string

	ProvinceID   string
	ProvinceName string

	WardID   string
	WardName string

	Type        AddressTypeEnum
	ContactName string
	PhoneNumber string
	AddressLine string
}

type NewShopParams struct {
	UserID shared.UserID
	Name   string
	Slug   string

	Description *string
	LogoUrl     *string
	BannerUrl   *string
}
