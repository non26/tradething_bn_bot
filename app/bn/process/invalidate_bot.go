package process

import (
	"context"
	"errors"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/process/domain"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
)

func (b *botService) InvalidateBot(ctx context.Context, req *domain.InvalidateBot) (*res.InvalidateBotHandlerResponse, error) {

	position := req.ToPosition()
	lookUpResult, err := b.lookUp.LookUp(ctx, position)
	if err != nil {
		return nil, err
	}

	if lookUpResult == nil {
		return nil, errors.New("bot order not found")
	}

	err = lookUpResult.ValidateBotOrderIDWith(req.BotOrderId)
	if err != nil {
		return nil, err
	}

	err = lookUpResult.ValiddatePositionSideWith(req.PositionSide)
	if err != nil {
		return nil, err
	}
	position.AmountB = lookUpResult.GetAmountB()
	if req.PositionSide == bnconstant.LONG {
		position.Side = bnconstant.SELL
	} else {
		position.Side = bnconstant.BUY
	}
	err = b.trade.InvalidatePosition(ctx, position)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
