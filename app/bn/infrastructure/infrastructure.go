package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure/position"
	domainservice "tradethingbot/app/bn/process/domain_service"
)

type IPositionBuilder interface {
	GetPosition(ctx context.Context, position_side string) position.IPosition
}

type ITrade interface {
	PlacePosition(ctx context.Context, position *position.Position) error
	InvalidatePosition(ctx context.Context, position *position.Position) error
}

type IBotLookUp interface {
	LookUp(ctx context.Context, position *position.Position) (result domainservice.ILookUpResult, err error)
}
