package module

import (
	"shop-command-module/internal/application/commands/add_shop_address"
	"shop-command-module/internal/application/commands/assign_member"
	"shop-command-module/internal/application/commands/create_shop"
	"shop-command-module/internal/application/queries/check_permission"
	"shop-command-module/internal/application/services/outbox"
	"shop-command-module/internal/application/services/shop"
	"shop-command-module/internal/domain/member"
)

type ApplicationModule struct {
	CreateShopExecutor      create_shop.Executor
	AssignMemberExecutor    assign_member.Executor
	AddShopAddressExecutor  add_shop_address.Executor
	CheckPermissionExecutor check_permission.Executor
}

func NewApplicationModule(infra *InfraModule) *ApplicationModule {
	outboxService := outbox.NewOutboxService(infra.OutboxRepo)
	authorizer := member.NewAuthorizer()

	shopService := shop.NewShopService(
		infra.ShopRepo,
		infra.MemberRepo,
		authorizer,
		outboxService,
		infra.ShopCache,
		infra.TxManager,
	)

	createShopCommandHandler := create_shop.NewHandler(shopService)
	assignMemberHandler := assign_member.NewHandler(shopService)
	addShopAddressHandler := add_shop_address.NewHandler(shopService)
	checkPermissionCommandHandler := check_permission.NewHandler(shopService)

	return &ApplicationModule{
		CreateShopExecutor:      createShopCommandHandler,
		AssignMemberExecutor:    assignMemberHandler,
		AddShopAddressExecutor:  addShopAddressHandler,
		CheckPermissionExecutor: checkPermissionCommandHandler,
	}
}
