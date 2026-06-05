package shop

import (
	"context"
	"shop-command-module/internal/domain/shared"
)

type QueryRepository interface {
	CheckSlugExists(ctx context.Context, slug string) (bool, error)
	FindByID(ctx context.Context, shopID shared.ShopID) (*Shop, error)
}

type CommandRepository interface {
	CreateShop(ctx context.Context, shop *Shop) error
	CreateAddress(ctx context.Context, addr *ShopAddress) error
	UpdateStatus(ctx context.Context, shop *Shop) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
