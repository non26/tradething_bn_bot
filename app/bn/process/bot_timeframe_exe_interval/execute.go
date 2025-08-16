package bottimeframeexeinterval

import (
	"context"
	"errors"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/process/domain"
)

func (b *botTimeframeExeInterval) Execute(ctx context.Context, req *domain.BotTimeframeExeIntervalRequest) (*res.BotTimeframeExeIntervalResponse, error) {

	// InTime := req.IsPresentInTimeframe()
	InTime := true

	lookUpResult, err := b.lookUp.LookUp(ctx, req.ToPosition())
	if err != nil {
		return nil, err
	}

	if lookUpResult == nil {
		return nil, errors.New("look up result not found")
	}

	if !lookUpResult.IsCurrentBotActive() {
		return nil, errors.New("current bot is not active")
	}

	req.SetSymbol(lookUpResult.GetSymbol())
	req.SetAccountId(lookUpResult.GetAccountId())
	req.SetIsActive(lookUpResult.IsCurrentBotActive())

	if !lookUpResult.IsFirstTime() {
		err = lookUpResult.ValidateBotOrderIDWith(req.GetBotOrderID())
		if err != nil {
			return nil, err
		}

		isPositionSideChanged := false
		err = lookUpResult.ValiddatePositionSideWith(req.GetPositionSide())
		if err != nil {
			isPositionSideChanged = true
		}

		isQuantityChanged := false
		err = lookUpResult.ValidateAmountQuantityWith(req.GetAmountB())
		if err != nil {
			isQuantityChanged = true
		}

		if isPositionSideChanged || isQuantityChanged {
			oldPosition := domain.BotTimeframeExeIntervalRequest{}
			oldPosition.SetAccountId(req.GetAccountId())
			oldPosition.SetBotId(req.GetBotId())
			oldPosition.SetBotOrderID(req.GetBotOrderID())
			oldPosition.SetSymbol(req.GetSymbol())
			oldPosition.SetPositionSide(lookUpResult.GetPositionSide())
			oldPosition.SetAmountB(lookUpResult.GetAmountB())

			err = b.trade.PlacePosition(ctx, oldPosition.ToClosePosition())
			if err != nil {
				return nil, err
			}

			lookUpResult.SetNewIsFirstTime(true)
			err = b.storeBotRegistor.Upsert(ctx, req.ToPosition())
			if err != nil {
				return nil, err
			}
		}
	}

	if InTime {
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
