package req

import (
	"tradethingbot/app/bn/process/domain"
)

type InvalidateBotHandlerRequest struct {
	BotId        string `json:"bot_id"`
	BotOrderId   string `json:"bot_order_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
}

func (b *InvalidateBotHandlerRequest) ToServiceModel() *domain.InvalidateBot {
	return &domain.InvalidateBot{
		BotId:        b.BotId,
		BotOrderId:   b.BotOrderId,
		Symbol:       b.Symbol,
		PositionSide: b.PositionSide,
	}
}
