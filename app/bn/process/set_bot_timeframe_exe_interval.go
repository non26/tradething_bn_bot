package process

import (
	"context"
	"errors"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/process/domain"
)

func (b *botService) SetBotTimeframeExeInterval(ctx context.Context, req *domain.BotTimeframeExeIntervalRequest) (*res.BotTimeframeExeIntervalDetailResponse, error) {
	lookUpResult, err := b.lookUp.LookUp(ctx, req.ToPosition())
	if err != nil {
		return nil, err
	}
	if lookUpResult != nil {
		return nil, errors.New("bot order already create")
	}
	// case when lookUpResult is nil

	err = b.store.Upsert(ctx, req.ToSetPosition())
	if err != nil {
		return nil, err
	}

	res := res.BotTimeframeExeIntervalDetailResponse{
		BotId:        req.GetBotId(),
		BotOrderID:   req.GetBotOrderID(),
		Symbol:       req.GetSymbol(),
		PositionSide: req.GetPositionSide(),
		AmountB:      req.GetAmountB(),
		IsActive:     false,
	}

	return &res, nil
}
