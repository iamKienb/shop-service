package shop

import (
	"context"
)

type QueryRepository interface {
	CheckSlugExists(ctx context.Context, slug string) (bool, error)
}

type CommandRepository interface {
	CreateShop(ctx context.Context, shop *Shop) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
