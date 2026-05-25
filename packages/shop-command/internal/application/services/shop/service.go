package shop

import (
	"context"
	"user-command-module/internal/application/commands/add_shop_address"
	"user-command-module/internal/application/commands/assign_member"
	"user-command-module/internal/application/commands/create_shop"
	"user-command-module/internal/application/port"
	"user-command-module/internal/application/queries/check_permission"
	"user-command-module/internal/application/services/outbox"
	"user-command-module/internal/domain/member"
	"user-command-module/internal/domain/shop"
)

type Service interface {
	CreateShop(ctx context.Context, cmd create_shop.Command) (*create_shop.Result, error)
	AssignMember(ctx context.Context, cmd assign_member.Command) (*assign_member.Result, error)
	AddAddress(ctx context.Context, cmd add_shop_address.Command) (*add_shop_address.Result, error)
	CheckPermission(ctx context.Context, qry check_permission.Query) (*check_permission.Result, error)
}

type shopService struct {
	shopRepo      shop.Repository
	memberRepo    member.Repository
	memAuthor     member.Authorizer
	outboxService outbox.Service
	shopCache     port.ShopCache
	txManager     port.TxManager
}

func NewShopService(
	shopRepo shop.Repository,
	memberRepo member.Repository,
	memAuthor member.Authorizer,
	outboxService outbox.Service,
	shopCache port.ShopCache,
	txManager port.TxManager,
) Service {
	return &shopService{
		shopRepo:      shopRepo,
		memberRepo:    memberRepo,
		memAuthor:     memAuthor,
		outboxService: outboxService,
		shopCache:     shopCache,
		txManager:     txManager,
	}
}
