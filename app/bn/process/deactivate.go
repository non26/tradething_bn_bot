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
		response.SetFailed("bot is already not active")
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

	currentPosition := domain.BotTimeframeExeIntervalRequest{}
	currentPosition.SetAccountId(bot.AccountId)
	currentPosition.SetBotId(bot.BotID)
	currentPosition.SetBotOrderID(bot.ClientId)
	currentPosition.SetSymbol(bot.Symbol)
	currentPosition.SetPositionSide(bot.PositionSide)
	currentPosition.SetAmountB(bot.AmountB)
	err = b.trade.PlacePosition(ctx, currentPosition.ToClosePosition())
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
		Setting:      bot.Setting,
	}

	err = b.storeBotRegistor.Upsert(ctx, deactivatePosition)
	if err != nil {
		response.SetFailed(err.Error())
		return response
	}

	err = b.storeBotOnRun.Delete(ctx, req.ToPosition())
	if err != nil {
		response.SetFailed(err.Error())
		return response
	}

	response.SetSuccess()
	return response
}
