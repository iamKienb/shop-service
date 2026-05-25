package handler

import (
	"context"
	"encoding/json"
	"shop-shared-module/events"
	"shop-worker-module/internal/application/port"
)

type ShopProfileCreatedHandler struct {
	repo  port.ESRepository
	alias string
}

func NewShopProfileCreatedHandler(repo port.ESRepository, alias string) *ShopProfileCreatedHandler {
	return &ShopProfileCreatedHandler{repo: repo, alias: alias}
}

func (h *ShopProfileCreatedHandler) Handle(ctx context.Context, raw json.RawMessage) error {
	var payload events.ShopProfileCreated
	if err := json.Unmarshal(raw, &payload); err != nil {
		return err
	}

	doc := map[string]any{
		"profile": map[string]any{
			"shop_id":     payload.ShopID,
			"description": payload.Description,
			"logo_url":    payload.LogoUrl,
			"banner_url":  payload.BannerUrl,
			"created_by":  payload.CreatedBy,
			"created_at":  payload.CreatedAt,
		},
	}

	return h.repo.SyncData(ctx, h.alias, payload.ShopID, doc)
}
