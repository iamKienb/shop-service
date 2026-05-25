package events

import "time"

const TopicShopProfileCreated = "shop-service.shop.profile.created"

type ShopProfileCreated struct {
	ShopID string `json:"shop_id"`

	Description *string `json:"description"`
	LogoUrl     *string `json:"logo_url"`
	BannerUrl   *string `json:"banner_url"`

	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
