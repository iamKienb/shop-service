package shop

import (
	"shop-command-module/internal/domain/shared"
	"time"
)

type AddressTypeEnum string

const (
	TypePickup AddressTypeEnum = "PICKUP"
	TypeReturn AddressTypeEnum = "RETURN"
)

type ShopAddress struct {
	ID     shared.ShopAddressID
	ShopID shared.ShopID

	CountryID   string
	CountryName string

	ProvinceID   string
	ProvinceName string

	WardID   string
	WardName string

	ContactName string
	PhoneNumber string
	AddressLine string
	Type        AddressTypeEnum

	CreatedBy shared.UserID
	UpdatedBy *shared.UserID

	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (e AddressTypeEnum) IsValidType() bool {
	switch e {
	case TypePickup, TypeReturn:
		return true
	}

	return false
}

func (e AddressTypeEnum) IsValid() bool {
	return e.IsValidType()
}
