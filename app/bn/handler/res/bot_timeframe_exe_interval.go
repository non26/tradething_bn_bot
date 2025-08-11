package res

type BotTimeframeExeIntervalResponse struct {
	BotOrderID string `json:"bot_order_id"`
	Message    string `json:"message"`
	Status     string `json:"status"`
	Code       string `json:"code"`
}

type BotTimeframeExeIntervalDetailResponse struct {
	BotId        string `json:"botId"`
	BotOrderID   string `json:"botOrderId"` // client id
	Symbol       string `json:"symbol"`
	PositionSide string `json:"positionSide"`
	AmountB      string `json:"amountB"`
	AccountId    string `json:"accountId"`
	IsActive     bool   `json:"isActive"`
}
