package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure"
	adaptor "tradethingbot/app/bn/infrastructure/adaptor/trade"
	"tradethingbot/app/bn/infrastructure/position"

	"github.com/non26/tradepkg/pkg/bn/utils"
)

type trade struct {
	position infrastructure.IPositionBuilder
	adaptor  adaptor.IOrderAdaptor
}

func NewTrade(
	position infrastructure.IPositionBuilder,
	adaptor adaptor.IOrderAdaptor,
) infrastructure.ITrade {
	return &trade{
		position: position,
		adaptor:  adaptor,
	}
}

func (t *trade) PlacePosition(ctx context.Context, position *position.Position) error {
	futurePosition := t.position.GetPosition(ctx, position.PositionSide)
	if utils.IsBuyPosition(position.Side, position.PositionSide) {
		return futurePosition.Buy(ctx, position)
	}
	return futurePosition.Sell(ctx, position)
}

func (t *trade) InvalidatePosition(ctx context.Context, position *position.Position) error {
	futurePosition := t.position.GetPosition(ctx, position.PositionSide)
	return futurePosition.Invalidate(ctx, position)
}
