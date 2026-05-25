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

func NewShopMemberAddedHandler(repo port.ESRepository, alias string) *ShopAddressAddedHandler {
	return &ShopAddressAddedHandler{repo: repo, alias: alias}
}

func (h *ShopMemberAddedHandler) Handle(ctx context.Context, raw json.RawMessage) error {
	var payload events.ShopMemberAdded
	if err := json.Unmarshal(raw, &payload); err != nil {
		return err
	}

	doc := map[string]any{
		"members": map[string]any{
			"member_id":   payload.MemberID,
			"member_name": payload.MemberName,

			"added_by":      payload.AddedBy,
			"name_added_by": payload.NameAddedBy,

			"joined_at": payload.JoinedAt,

			// "roles": map[string]any{
			// 	"id": payload.Roles
			// 	"code"
			// 	"name"
			// }
		},
	}

	return h.repo.SyncData(ctx, h.alias, payload.ShopID, doc)
}
