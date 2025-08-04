package infrastructure

import (
	"context"
	adaptor "tradethingbot/app/bn/infrastructure/adaptor/trade"
	"tradethingbot/app/bn/infrastructure/position"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

type trade struct {
	position IPositionBuilder
	adaptor  adaptor.IOrderAdaptor
}

func NewTrade(
	position IPositionBuilder,
	adaptor adaptor.IOrderAdaptor,
) ITrade {
	return &trade{
		position: position,
		adaptor:  adaptor,
	}
}

func (t *trade) PlacePosition(ctx context.Context, position *position.Position) error {
	futurePosition := t.position.GetPosition(ctx, position.PositionSide)
	if position.Side == bnconstant.BUY {
		return futurePosition.Buy(ctx, position)
	}
	return futurePosition.Sell(ctx, position)
}

func (t *trade) InvalidatePosition(ctx context.Context, position *position.Position) error {
	futurePosition := t.position.GetPosition(ctx, position.PositionSide)
	return futurePosition.Invalidate(ctx, position)
}
