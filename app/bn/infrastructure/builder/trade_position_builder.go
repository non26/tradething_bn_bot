package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type positionBuilder struct {
	longPosition  infrastructure.IPosition
	shortPosition infrastructure.IPosition
}

func NewFuturePosition(longPosition infrastructure.IPosition, shortPosition infrastructure.IPosition) infrastructure.IPositionBuilder {
	return &positionBuilder{
		longPosition:  longPosition,
		shortPosition: shortPosition,
	}
}

func (f *positionBuilder) GetPosition(ctx context.Context, position_side string) infrastructure.IPosition {
	if position_side == bnconstant.LONG {
		return f.longPosition
	}
	return f.shortPosition
}
