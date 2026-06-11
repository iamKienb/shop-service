package handler

import (
	"context"
	"encoding/json"
	"shop-shared-module/events"
	"shop-worker-module/internal/application/port"
)

type ShopMemberAddedHandler struct {
	repo  port.ESRepository
	alias string
}

func NewShopMemberAddedHandler(repo port.ESRepository, alias string) *ShopMemberAddedHandler {
	return &ShopMemberAddedHandler{repo: repo, alias: alias}
}

func (h *ShopMemberAddedHandler) Handle(ctx context.Context, raw json.RawMessage) error {
	var payload events.ShopMemberAdded
	if err := json.Unmarshal(raw, &payload); err != nil {
		return err
	}

	roleIDs := make([]int32, 0, len(payload.Roles))
	for _, role := range payload.Roles {
		roleIDs = append(roleIDs, int32(role.ID))
	}

	doc := map[string]any{
		"id":       payload.MemberID,
		"name":     payload.MemberName,
		"role_ids": roleIDs,

		"added_by":      payload.AddedBy,
		"name_added_by": payload.NameAddedBy,

		"joined_at": payload.JoinedAt,
		"roles":     payload.Roles,
	}

	return h.repo.SyncNestedData(ctx, port.NestedParam{
		Index:         h.alias,
		DocID:         payload.ShopID,
		NestedField:   "members",
		NestedFieldID: payload.MemberID,
		Data:          doc,
	})
}
