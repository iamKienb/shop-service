package events

import "time"

const TopicShopAddressAdded = "shop-service.shop.address.added"

type ShopAddressAdded struct {
	ShopID        string `json:"shop_id"`
	ShopAddressID string `json:"shop_address_id"`

	CountryID   string `json:"country_id"`
	CountryName string `json:"country_name"`

	ProvinceID   string `json:"province_id"`
	ProvinceName string `json:"province_name"`

	WardID   string `json:"ward_id"`
	WardName string `json:"ward_name"`

	AddressLine  string `json:"address_line"`
	ContractName string `json:"contract_name"`
	PhoneNumber  string `json:"phone_number"`
	Type         string `json:"type"`

	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
