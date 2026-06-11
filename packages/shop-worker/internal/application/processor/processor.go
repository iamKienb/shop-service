package processor

import (
	"context"
	"fmt"
	"shop-shared-module/alias"
	"shop-shared-module/events"
	"shop-worker-module/internal/application/port"
	"shop-worker-module/internal/application/processor/handler"
	"time"
)

const (
	idemKeyTTL = 24 * time.Hour
	key        = "shop-worker:key:%s"
)

type ShopEventProcessor struct {
	handlers    map[string]port.EventHandler
	workerCache port.WorkerCache
}

func NewShopEventProcessor(repo port.ESRepository, workerCache port.WorkerCache) port.EventProcessor {
	shopAlias := alias.ShopAlias

	p := &ShopEventProcessor{
		handlers:    make(map[string]port.EventHandler),
		workerCache: workerCache,
	}

	p.handlers[events.TopicShopCreated] = handler.NewShopCreatedHandler(repo, shopAlias)
	p.handlers[events.TopicShopProfileUpdated] = handler.NewShopProfileUpdatedHandler(repo, shopAlias)
	p.handlers[events.TopicShopAddressAdded] = handler.NewShopAddressAddedHandler(repo, shopAlias)
	p.handlers[events.TopicShopMemberAdded] = handler.NewShopMemberAddedHandler(repo, shopAlias)

	return p
}

func (p *ShopEventProcessor) Handle(ctx context.Context, msg port.Message) error {
	h, ok := p.handlers[msg.Topic]
	if !ok {
		return nil
	}
	idemKey := msg.IdempotencyKey()

	if idemKey != "" {
		key := fmt.Sprintf(key, idemKey)
		isNew, err := p.workerCache.SetNx(ctx, key, 1, idemKeyTTL)
		if err != nil {
			return err
		}
		if !isNew {
			return nil
		}
	}

	return h.Handle(ctx, msg.Value)
}
