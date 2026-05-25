package processor

import (
	"context"
	"encoding/json"
	"shop-shared-module/alias"
	"shop-shared-module/events"
	"shop-worker-module/internal/application/port"
	"shop-worker-module/internal/application/processor/handler"
)

type ShopEventProcessor struct {
	handlers map[string]port.EventHandler
}

func NewShopEventProcessor(repo port.ESRepository) port.EventProcessor {
	shopAlias := alias.ShopAlias

	p := &ShopEventProcessor{
		handlers: make(map[string]port.EventHandler),
	}

	p.handlers[events.TopicShopCreated] = handler.NewShopCreatedHandler(repo, shopAlias)
	p.handlers[events.TopicShopProfileCreated] = handler.NewShopProfileCreatedHandler(repo, shopAlias)
	p.handlers[events.TopicShopAddressAdded] = handler.NewShopAddressAddedHandler(repo, shopAlias)
	p.handlers[events.TopicShopMemberAdded] = handler.NewShopMemberAddedHandler(repo, shopAlias)

	return p
}

func (p *ShopEventProcessor) Handle(ctx context.Context, msg port.Message) error {
	h, ok := p.handlers[msg.Topic]
	if !ok {
		return nil
	}

	var rawPayload json.RawMessage = msg.Value

	return h.Handle(ctx, rawPayload)
}
