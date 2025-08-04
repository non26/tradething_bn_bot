package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure/position"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type positionBuilder struct {
	longPosition  position.IPosition
	shortPosition position.IPosition
}

func NewFuturePosition(longPosition position.IPosition, shortPosition position.IPosition) IPositionBuilder {
	return &positionBuilder{
		longPosition:  longPosition,
		shortPosition: shortPosition,
	}
}

func (f *positionBuilder) GetPosition(ctx context.Context, position_side string) position.IPosition {
	if position_side == bnconstant.LONG {
		return f.longPosition
	}
	return f.shortPosition
}
