package port

import "context"

type NestedParam struct {
	Index         string
	DocID         string
	NestedField   string
	NestedFieldID string
	Data          any
}
type ESRepository interface {
	SyncData(ctx context.Context, index string, id string, data any) error
	SyncNestedData(ctx context.Context, param NestedParam) error
}
