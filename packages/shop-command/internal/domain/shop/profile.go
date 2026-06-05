package shop

import (
	"shop-command-module/internal/domain/shared"
	"time"
)

type Profile struct {
	ShopID      shared.ShopID
	Description *string
	LogoUrl     *string
	BannerUrl   *string

	CreatedBy shared.UserID
	UpdatedBy *shared.UserID

	CreatedAt time.Time
	UpdatedAt *time.Time
}
