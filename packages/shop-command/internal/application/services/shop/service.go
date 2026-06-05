package shop

import (
	"context"
	"shop-command-module/internal/application/commands/add_shop_address"
	"shop-command-module/internal/application/commands/assign_member"
	"shop-command-module/internal/application/commands/create_shop"
	"shop-command-module/internal/application/port"
	"shop-command-module/internal/application/queries/check_permission"
	"shop-command-module/internal/application/services/outbox"
	"shop-command-module/internal/domain/member"
	"shop-command-module/internal/domain/shop"
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
