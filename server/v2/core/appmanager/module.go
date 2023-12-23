package appmanager

import (
	"context"

	"cosmossdk.io/server/v2/core/transaction"
)

type Identity = []byte

type MsgRouterBuilder interface {
	RegisterHandler(msg Type, handlerFunc func(ctx context.Context, msg Type) (resp Type, err error))
}

type QueryRouterBuilder = MsgRouterBuilder

type PreMsgRouterBuilder interface {
	RegisterPreHandler(msg Type, preHandler func(ctx context.Context, msg Type) error)
}

type PostMsgRouterBuilder interface {
	RegisterPostHandler(msg Type, postHandler func(ctx context.Context, msg, msgResp Type) error)
}

type STFModule[T transaction.Tx] interface {
	Name() string
	RegisterMsgHandlers(router MsgRouterBuilder)
	RegisterQueryHandler(router QueryRouterBuilder)
	BeginBlocker() func(ctx context.Context) error
	EndBlocker() func(ctx context.Context) error
	UpdateValidators() func(ctx context.Context) ([]ValidatorUpdate, error)
	TxValidator() func(ctx context.Context, tx T) error // why does the module handle registration
	RegisterPreMsgHandler(router PreMsgRouterBuilder)
	RegisterPostMsgHandler(router PostMsgRouterBuilder)
}

type Module[T transaction.Tx] interface {
	STFModule[T]
}

// Update defines what is expected to be returned
type ValidatorUpdate struct {
	PubKey []byte
	Power  int64 // updated power of the validtor
}
