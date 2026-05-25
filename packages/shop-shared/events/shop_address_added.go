package events

import "time"

const TopicShopAddressAdded = "shop-service.shop.address.added"

type ShopAddressAdded struct {
	ShopID        string `json:"shop_id"`
	ShopAddressID string `json:"shop_address_id"`

	CountryID   int    `json:"country_id"`
	CountryName string `json:"country_name"`

	CityID   int    `json:"city_id"`
	CityName string `json:"city_name"`

	DistrictID   int    `json:"district_id"`
	DistrictName string `json:"district_name"`

	WardID   int    `json:"ward_id"`
	WardName string `json:"ward_name"`

	AddressLine  string `json:"address_line"`
	ContractName string `json:"contract_name"`
	PhoneNumber  string `json:"phone_number"`
	Type         string `json:"type"`

	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
