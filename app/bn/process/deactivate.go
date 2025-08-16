package process

import (
	"context"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/process/domain"
)

func (b *botService) DeactivateBot(ctx context.Context, req *domain.Activation) *res.ActivationResponse {

	response := &res.ActivationResponse{
		BotId:      req.BotId,
		BotOrderId: req.BotOrderId,
	}
	lookUpResult, err := b.lookUp.LookUp(ctx, req.ToPosition())
	if err != nil {
		response.SetFailed(err.Error())
		return response
	}
	if !lookUpResult.IsCurrentBotActive() {
		response.SetFailed("bot is not active")
		return response
	}
	bot, err := b.storeBotOnRun.Get(ctx, req.ToPosition())
	if err != nil {
		response.SetFailed(err.Error())
		return response
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
		return response
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
		return response
	}

	return response
}
