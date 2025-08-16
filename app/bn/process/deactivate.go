package process

import (
	"context"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/process/domain"
)

func (b *botService) DeactivateBot(ctx context.Context, req []domain.Activation) ([]res.ActivationResponse, error) {
	responses := make([]res.ActivationResponse, len(req))
	for _, deactivation := range req {
		response := res.ActivationResponse{
			BotId:      deactivation.BotId,
			BotOrderId: deactivation.BotOrderId,
		}
		lookUpResult, err := b.lookUp.LookUp(ctx, deactivation.ToPosition())
		if err != nil {
			return nil, err
		}
		if !lookUpResult.IsCurrentBotActive() {
			response.SetFailed("bot is not active")
			responses = append(responses, response)
			continue
		}
		bot, err := b.storeBotOnRun.Get(ctx, deactivation.ToPosition())
		if err != nil {
			response.SetFailed(err.Error())
			responses = append(responses, response)
			continue
		}

		closePosition := domain.BotTimeframeExeIntervalRequest{}
		closePosition.SetAccountId(bot.AccountId)
		closePosition.SetBotId(bot.BotID)
		closePosition.SetBotOrderID(bot.ClientId)
		closePosition.SetSymbol(bot.Symbol)
		closePosition.SetPositionSide(bot.PositionSide)
		closePosition.SetAmountB(bot.AmountB)
		err = b.trade.PlacePosition(ctx, closePosition.ToClosePosition())
		if err != nil {
			response.SetFailed(err.Error())
			responses = append(responses, response)
			continue
		}

		deactivatePosition := &infrastructure.Position{
			BotID:        bot.BotID,
			ClientId:     bot.ClientId,
			IsActive:     false,
			AmountB:      bot.AmountB,
			Symbol:       bot.Symbol,
			PositionSide: bot.PositionSide,
			AccountId:    bot.AccountId,
		}

		err = b.storeBotOnRun.Upsert(ctx, deactivatePosition)
		if err != nil {
			response.SetFailed(err.Error())
			responses = append(responses, response)
			continue
		}
	}
	return responses, nil
}
