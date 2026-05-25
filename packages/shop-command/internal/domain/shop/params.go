package shop

import "user-command-module/internal/domain/shared"

type NewShopAddressParams struct {
	UserID shared.UserID
	ShopID shared.ShopID

	CountryID   int
	CountryName string

	CityID   int
	CityName string

	DistrictID   int
	DistrictName string

	WardID   int
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
	Address     NewShopAddressParams
}
