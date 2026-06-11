package handler

import (
	"context"
	"encoding/json"
	"shop-shared-module/events"
	"shop-worker-module/internal/application/port"
)

type ShopCreatedHandler struct {
	repo  port.ESRepository
	alias string
}

func NewShopCreatedHandler(repo port.ESRepository, alias string) *ShopCreatedHandler {
	return &ShopCreatedHandler{repo: repo, alias: alias}
}

func (h *ShopCreatedHandler) Handle(ctx context.Context, raw json.RawMessage) error {
	var payload events.ShopCreated
	if err := json.Unmarshal(raw, &payload); err != nil {
		return err
	}

	doc := map[string]any{
		"id":         payload.ShopID,
		"owner_id":   payload.OwnerID,
		"name":       payload.Name,
		"slug":       payload.Slug,
		"status":     payload.Status,
		"profile":    payload.Profile,
		"created_by": payload.CreatedBy,
		"created_at": payload.CreatedAt,
	}

	return h.repo.SyncData(ctx, h.alias, payload.ShopID, doc)
}
