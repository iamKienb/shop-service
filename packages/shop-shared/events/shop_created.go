package events

import "time"

const TopicShopCreated = "shop-service.shop.created"

type ShopCreated struct {
	ShopID  string  `json:"shop_id"`
	OwnerID string  `json:"owner_id"`
	Name    string  `json:"name"`
	Slug    string  `json:"slug"`
	Status  string  `json:"status"`
	Profile Profile `json:"profile"`

	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

type Profile struct {
	ShopID      string    `json:"shop_id"`
	Description *string   `json:"description"`
	LogoURL     *string   `json:"logo_url"`
	BannerURL   *string   `json:"banner_url"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}
