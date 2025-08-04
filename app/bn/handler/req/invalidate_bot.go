package req

import (
	"tradethingbot/app/bn/process/domain"
)

type InvalidateBotHandlerRequest struct {
	BotId      string `json:"botId"`
	BotOrderId string `json:"botOrderId"`
	AccountId  string `json:"accountId"`
}

func (b *InvalidateBotHandlerRequest) ToServiceModel() *domain.InvalidateBot {
	return &domain.InvalidateBot{
		BotId:      b.BotId,
		BotOrderId: b.BotOrderId,
		AccountId:  b.AccountId,
	}
}
