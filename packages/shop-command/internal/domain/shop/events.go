package shop

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type ShopCreatedEvent struct {
	ShopID    shared.ShopID
	OwnerID   shared.UserID
	Name      string
	Slug      string
	Status    ShopStatus
	CreatedBy shared.UserID
	CreatedAt time.Time
}

func (e ShopCreatedEvent) EventName() string {
	return "shop-service.shop.created"
}

func (e ShopCreatedEvent) IntegrationPayload() map[string]interface{} {
	return map[string]interface{}{
		"shop_id":    e.ShopID.String(),
		"owner_id":   e.OwnerID.String(),
		"name":       e.Name,
		"slug":       e.Slug,
		"status":     string(e.Status),
		"created_by": e.CreatedBy.String(),
		"created_at": e.CreatedAt,
	}
}

type ShopProfileCreatedEvent struct {
	ShopID      shared.ShopID
	Description *string
	LogoUrl     *string
	BannerUrl   *string
	CreatedBy   shared.UserID
	CreatedAt   time.Time
}

func (e ShopProfileCreatedEvent) EventName() string {
	return "shop-service.shop.created"
}

func (e ShopProfileCreatedEvent) IntegrationPayload() map[string]interface{} {
	return map[string]interface{}{
		"shop_id":     e.ShopID.String(),
		"description": valueOrNil(e.Description),
		"logo_url":    valueOrNil(e.LogoUrl),
		"banner_url":  valueOrNil(e.BannerUrl),
		"created_by":  e.CreatedBy.String(),
		"created_at":  e.CreatedAt,
	}
}

type ShopActivatedEvent struct {
	ShopID    shared.ShopID
	Status    ShopStatus
	UpdatedBy shared.UserID
	UpdatedAt time.Time
}

func (e ShopActivatedEvent) EventName() string {
	return "shop-service.shop.activated"
}

func (e ShopActivatedEvent) IntegrationPayload() map[string]interface{} {
	return map[string]interface{}{
		"shop_id":    e.ShopID.String(),
		"status":     e.Status,
		"updated_by": e.UpdatedBy.String(),
		"updated_at": e.UpdatedAt,
	}
}

func valueOrNil(value *string) any {
	if value == nil {
		return nil
	}

	return *value
}

type ShopAddressAddedEvent struct {
	ShopID        shared.ShopID
	ShopAddressID shared.ShopAddressID

	CountryID   int
	CountryName string

	CityID   int
	CityName string

	DistrictID   int
	DistrictName string

	WardID   int
	WardName string

	AddressLine string
	ContactName string
	PhoneNumber string

	Type AddressTypeEnum

	CreatedBy shared.UserID
	CreatedAt time.Time
}

func (e ShopAddressAddedEvent) EventName() string {
	return "shop-service.shop.address.added"
}

func (e ShopAddressAddedEvent) IntegrationPayload() map[string]interface{} {
	return map[string]interface{}{
		"shop_id":         e.ShopID.String(),
		"shop_address_id": e.ShopAddressID.String(),
		"country_id":      e.CountryID,
		"country_name":    e.CountryName,
		"city_id":         e.CityID,
		"city_name":       e.CityName,
		"district_id":     e.DistrictID,
		"district_name":   e.DistrictName,
		"ward_id":         e.WardID,
		"ward_name":       e.WardName,
		"address_line":    e.AddressLine,
		"contact_name":    e.ContactName,
		"phone_number":    e.PhoneNumber,
		"type":            string(e.Type),
		"created_by":      e.CreatedBy.String(),
		"created_at":      e.CreatedAt,
	}
}
