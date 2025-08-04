package domain

import "tradethingbot/app/bn/infrastructure/position"

type InvalidateBot struct {
	BotId        string `json:"bot_id"`
	BotOrderId   string `json:"bot_order_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
}

func (b *InvalidateBot) ToPosition() *position.Position {
	return &position.Position{
		BotID:        b.BotId,
		ClientId:     b.BotOrderId,
		Symbol:       b.Symbol,
		PositionSide: b.PositionSide,
	}
}
