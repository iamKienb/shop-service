package module

import (
	"log/slog"
	"net/http"

	shopAdapter "shop-query-module/internal/adapter/shop"

	"connectrpc.com/grpcreflect"
	"github.com/iamKienb/api-contract/gen/shop/shopconnect"
	observabilityx "github.com/iamKienb/go-core/middleware/observability"
)

type AdapterModule struct {
	Mux *http.ServeMux
}

func NewAdapterModule(app *ApplicationModule, logger *slog.Logger) *AdapterModule {
	server := shopAdapter.NewQueryServer(
		app.GetShopDetailExecutor,
		app.SearchShopsExecutor,
		app.ListShopAddressesExecutor,
		app.ListShopMembersExecutor,
	)

	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(shopconnect.ShopQueryServiceName)
	mux.Handle(shopconnect.NewShopQueryServiceHandler(server, observabilityx.InternalServerOption(logger)))
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	return &AdapterModule{Mux: mux}
}
