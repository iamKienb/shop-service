package events

import "time"

const TopicShopCreated = "shop-service.shop.created"

type ShopCreated struct {
	ShopID  string `json:"shop_id"`
	OwnerID string `json:"owner_id"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Status  string `json:"status"`

	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
