package handler

import (
	"context"
	"encoding/json"
	"shop-shared-module/events"
	"shop-worker-module/internal/application/port"
)

type ShopAddressAddedHandler struct {
	repo  port.ESRepository
	alias string
}

func NewShopAddressAddedHandler(repo port.ESRepository, alias string) *ShopAddressAddedHandler {
	return &ShopAddressAddedHandler{repo: repo, alias: alias}
}

func (h *ShopAddressAddedHandler) Handle(ctx context.Context, raw json.RawMessage) error {
	var payload events.ShopAddressAdded
	if err := json.Unmarshal(raw, &payload); err != nil {
		return err
	}

	fullAddress := payload.AddressLine + payload.WardName + payload.DistrictName + payload.CityName

	doc := map[string]any{
		"address": map[string]any{
			"id":      payload.ShopAddressID,
			"shop_id": payload.ShopID,

			"full_address": fullAddress,
			"country": map[string]any{
				"country_id": payload.CountryID,
				"name":       payload.CountryName,
			},
			"city": map[string]any{
				"city_id": payload.CityID,
				"name":    payload.CityName,
			},
			"district": map[string]any{
				"district_id": payload.DistrictID,
				"name":        payload.DistrictName,
			},
			"ward": map[string]any{
				"ward_id": payload.WardID,
				"name":    payload.WardName,
			},

			"address_line": payload.AddressLine,
			"contact_name": payload.ContractName,
			"phone_number": payload.PhoneNumber,
			"type":         payload.Type,
			"created_by":   payload.CreatedBy,
			"created_at":   payload.CreatedAt,
		},
	}

	return h.repo.SyncData(ctx, h.alias, payload.ShopID, doc)
}
