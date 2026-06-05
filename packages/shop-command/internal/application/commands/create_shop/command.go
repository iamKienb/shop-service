package create_shop

import (
	"context"
	"shop-command-module/internal/domain/shared"
)

type User struct {
	ID   shared.UserID
	Name string
}

type Profile struct {
	Description *string
	LogoUrl     *string
	BannerUrl   *string
}

type Command struct {
	User    User
	Name    string
	Slug    string
	Profile *Profile
}

type Result struct {
	ShopID string
}

type Executor interface {
	Execute(ctx context.Context, cmd Command) (*Result, error)
}
