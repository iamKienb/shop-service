package shop

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type AddressTypeEnum string

const (
	TypePickup AddressTypeEnum = "PICKUP"
	TypeReturn AddressTypeEnum = "RETURN"
)

type ShopAddress struct {
	ID     shared.ShopAddressID
	ShopID shared.ShopID

	CountryID  int
	CityID     int
	DistrictID int
	WardID     int

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
