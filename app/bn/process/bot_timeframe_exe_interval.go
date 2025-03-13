package process

import (
	"context"
	"errors"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/process/domain"
)

func (b *botService) BotTimeframeExeInterval(ctx context.Context, req *domain.BotTimeframeExeIntervalRequest) (*res.BotTimeframeExeIntervalResponse, error) {

	InTime := req.IsPresentInTimeframe()

	lookUpResult, err := b.lookUp.LookUp(ctx, req.ToPosition())
	if err != nil {
		return nil, err
	}

	if lookUpResult == nil {
		return nil, errors.New("bot order not found")
	}

	err = lookUpResult.ValidateBotOrderIDWith(req.GetBotOrderID())
	if err != nil {
		return nil, err
	}

	err = lookUpResult.ValiddatePositionSideWith(req.GetPositionSide())
	if err != nil {
		return nil, err
	}

	if InTime {
		if !lookUpResult.IsCurrentBotActive() {
			return nil, errors.New("current bot is not active")
		}

		if !lookUpResult.IsFirstTime() {
			err = b.trade.PlacePosition(ctx, req.ToClosePosition())
			if err != nil {
				return nil, err
			}
		}

		err = b.trade.PlacePosition(ctx, req.ToOpenPosition())
		if err != nil {
			return nil, err
		}
	} else {
		err = b.trade.InvalidatePosition(ctx, req.ToClosePosition())
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}
