package infrastructure

import (
	"context"
	domainservice "tradethingbot/app/bn/process/domain_service"
)

type IPositionBuilder interface {
	GetPosition(ctx context.Context, position_side string) IPosition
}

type ITrade interface {
	PlacePosition(ctx context.Context, position *Position) error
	InvalidatePosition(ctx context.Context, position *Position) error
}

type IBotLookUp interface {
	LookUp(ctx context.Context, position *Position) (result domainservice.ILookUpResult, err error)
	// LookUpByBotIdAndBotOrderId(ctx context.Context, botId string, botOrderId string) (result domainservice.ILookUpResult, err error)
}

type IBotOnRunStore interface {
	Upsert(ctx context.Context, position *Position) error
	GetAll(ctx context.Context) ([]*Position, error)
	Get(ctx context.Context, position *Position) (*Position, error)
	Delete(ctx context.Context, position *Position) error
}

type IBotRegistorStore interface {
	Upsert(ctx context.Context, position *Position) error
	Get(ctx context.Context, position *Position) (*Position, error)
	GetAll(ctx context.Context) ([]*Position, error)
	Delete(ctx context.Context, position *Position) error
}

type IBnFutureHistoryStore interface {
	Get(ctx context.Context, position *Position) (*Position, error)
	Insert(ctx context.Context, position *Position) error
}

type IBotStore interface {
	Get(ctx context.Context, botId string) (*Position, error)
}
