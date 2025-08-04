package domain

import "tradethingbot/app/bn/infrastructure/position"

type InvalidateBot struct {
	BotId      string `json:"botId"`
	BotOrderId string `json:"botOrderId"`
	AccountId  string `json:"accountId"`
	// Symbol       string `json:"symbol"`
	// PositionSide string `json:"position_side"`
}

func (b *InvalidateBot) ToPosition() *position.Position {
	return &position.Position{
		BotID:     b.BotId,
		ClientId:  b.BotOrderId,
		AccountId: b.AccountId,
		// Symbol:       b.Symbol,
		// PositionSide: b.PositionSide,
	}
}
