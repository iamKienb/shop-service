package module

import (
	"log/slog"
	"net/http"

	"shop-command-module/internal/adapter/shop"

	"connectrpc.com/grpcreflect"
	"github.com/iamKienb/api-contract/gen/shop/shopconnect"
	observabilityx "github.com/iamKienb/go-core/middleware/observability"
)

type AdapterModule struct {
	Mux *http.ServeMux
}

func NewAdapterModule(app *ApplicationModule, logger *slog.Logger) *AdapterModule {
	allInterceptors := observabilityx.InternalServerOption(logger)

	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		shopconnect.ShopCommandServiceName,
	)

	shopServer := shop.NewShopServer(
		app.CreateShopExecutor,
		app.AssignMemberExecutor,
		app.AddShopAddressExecutor,
		app.CheckPermissionExecutor,
	)

	mux.Handle(shopconnect.NewShopCommandServiceHandler(shopServer, allInterceptors))

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	return &AdapterModule{Mux: mux}
}
