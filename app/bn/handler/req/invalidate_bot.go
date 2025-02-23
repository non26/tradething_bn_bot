package req

type InvalidateBotHandlerRequest struct {
	BotId        string `json:"bot_id"`
	BotOrderId   string `json:"bot_order_id"`
	Symbol       string `json:"symbol"`
	PositionSide string `json:"position_side"`
}

// func (b *InvalidateBotHandlerRequest) ToServiceModel() *bnbotsvcreq.InvalidateBot {
// 	return &bnbotsvcreq.InvalidateBot{
// 		BotId:        b.BotId,
// 		BotOrderId:   b.BotOrderId,
// 		Symbol:       b.Symbol,
// 		PositionSide: b.PositionSide,
// 	}
// }
