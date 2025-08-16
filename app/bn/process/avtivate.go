package process

import (
	"context"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/process/domain"
)

func (b *botService) ActivateBot(ctx context.Context, req []domain.Activation) ([]res.ActivationResponse, error) {
	responses := make([]res.ActivationResponse, 0)
	for _, activation := range req {
		response := res.ActivationResponse{
			BotId:      activation.BotId,
			BotOrderId: activation.BotOrderId,
		}
		lookUpResult, err := b.lookUp.LookUp(ctx, activation.ToPosition())
		if err != nil {
			return nil, err
		}
		if lookUpResult.IsCurrentBotActive() {
			response.SetFailed("bot already active")
			responses = append(responses, response)
			continue
		}

		bot, err := b.storeBotOnRun.Get(ctx, activation.ToPosition())
		if err != nil {
			response.SetFailed(err.Error())
			responses = append(responses, response)
			continue
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
			responses = append(responses, response)
			continue
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
			responses = append(responses, response)
			continue
		}
		response.SetSuccess()
		responses = append(responses, response)
	}

	return responses, nil
}
