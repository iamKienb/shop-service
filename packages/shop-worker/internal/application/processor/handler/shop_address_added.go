package handler

import (
	"context"
	"encoding/json"
	"shop-shared-module/events"
	"shop-worker-module/internal/application/port"
	"strings"
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

	fullAddress := strings.Join([]string{
		payload.AddressLine,
		payload.WardName,
		payload.ProvinceName,
		payload.CountryName,
	}, ", ")

	doc := map[string]any{
		"id":      payload.ShopAddressID,
		"shop_id": payload.ShopID,

		"country": map[string]any{
			"id":   payload.CountryID,
			"name": payload.CountryName,
		},
		"province": map[string]any{
			"id":   payload.ProvinceID,
			"name": payload.ProvinceName,
		},
		"ward": map[string]any{
			"id":   payload.WardID,
			"name": payload.WardName,
		},

		"address_line": payload.AddressLine,
		"full_address": fullAddress,

		"contact_name": payload.ContractName,
		"phone_number": payload.PhoneNumber,
		"type":         payload.Type,
		"created_by":   payload.CreatedBy,
		"created_at":   payload.CreatedAt,
	}

	return h.repo.SyncNestedData(ctx, port.NestedParam{
		Index:         h.alias,
		DocID:         payload.ShopID,
		NestedField:   "addresses",
		NestedFieldID: payload.ShopAddressID,
		Data:          doc,
	})
}
