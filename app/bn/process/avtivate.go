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

	if !lookUpResult.IsRegistor() {
		response.SetFailed("bot registor not found")
		return response
	}

	bot, err := b.storeBotRegistor.Get(ctx, req.ToPosition())
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

	err = b.storeBotRegistor.Upsert(ctx, activatePosition)
	if err != nil {
		response.SetFailed(err.Error())
		return response
	}

	response.SetSuccess()
	return response
}
