package process

import (
	"context"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/process/domain"
)

func (b *botService) ActivateBot(ctx context.Context, req *domain.Activation) *res.ActivationResponse {
	response := &res.ActivationResponse{
		BotId:      req.BotId,
		BotOrderId: req.BotOrderId,
	}
	lookUpResult, err := b.lookUp.LookUp(ctx, req.ToPosition())
	if err != nil {
		response.SetFailed(err.Error())
		return response
	}
	if lookUpResult.IsCurrentBotActive() {
		response.SetFailed("bot already active")
		return response
	}

	bot, err := b.storeBotOnRun.Get(ctx, req.ToPosition())
	if err != nil {
		response.SetFailed(err.Error())
		return response
	}

	activatePosition := &infrastructure.Position{
		BotID:        bot.BotID,
		ClientId:     bot.ClientId,
		IsActive:     true,
		AmountB:      bot.AmountB,
		Symbol:       bot.Symbol,
		PositionSide: bot.PositionSide,
		AccountId:    bot.AccountId,
	}

	err = b.storeBotOnRun.Upsert(ctx, activatePosition)
	if err != nil {
		response.SetFailed(err.Error())
		return response
	}

	tfIntervalRequest := &domain.BotTimeframeExeIntervalRequest{}
	tfIntervalRequest.SetBotId(bot.BotID)
	tfIntervalRequest.SetBotOrderID(bot.ClientId)
	tfIntervalRequest.SetSymbol(bot.Symbol)
	tfIntervalRequest.SetPositionSide(bot.PositionSide)
	tfIntervalRequest.SetAmountB(bot.AmountB)
	tfIntervalRequest.SetAccountId(bot.AccountId)
	_, err = b.BotTimeframeExeInterval(ctx, tfIntervalRequest)
	if err != nil {
		response.SetFailed(err.Error())
		return response
	}
	response.SetSuccess()

	return response
}
