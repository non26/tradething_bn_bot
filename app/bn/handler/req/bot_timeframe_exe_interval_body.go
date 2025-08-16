package req

import (
	"tradethingbot/app/bn/process/domain"
)

type BotTimeframeExeIntervalHandlerRequest struct {
	BotId        string `json:"botId"`
	BotOrderID   string `json:"botOrderId"` // client id
	PositionSide string `json:"positionSide"`
	AmountB      string `json:"amountB"`
}

func (b *BotTimeframeExeIntervalHandlerRequest) ToBotServiceRequest() *domain.BotTimeframeExeIntervalRequest {
	svcmodel := &domain.BotTimeframeExeIntervalRequest{}
	svcmodel.SetBotId(b.BotId)
	svcmodel.SetBotOrderID(b.BotOrderID)
	svcmodel.SetPositionSide(b.PositionSide)
	svcmodel.SetAmountB(b.AmountB)
	return svcmodel
}

type SetBotTimeframeExeIntervalHandlerRequest struct {
	BotId        string `json:"botId"`
	BotOrderID   string `json:"botOrderId"` // client id
	PositionSide string `json:"positionSide"`
	AmountB      string `json:"amountB"`
	IsActive     bool   `json:"isActive,omitempty"`
	Symbol       string `json:"symbol"`
	AccountId    string `json:"accountId"`
}

func (b *SetBotTimeframeExeIntervalHandlerRequest) ToBotServiceRequest() *domain.BotTimeframeExeIntervalRequest {
	svcmodel := &domain.BotTimeframeExeIntervalRequest{}
	svcmodel.SetBotId(b.BotId)
	svcmodel.SetBotOrderID(b.BotOrderID)
	svcmodel.SetPositionSide(b.PositionSide)
	svcmodel.SetAmountB(b.AmountB)
	svcmodel.SetIsActive(b.IsActive)
	svcmodel.SetSymbol(b.Symbol)
	svcmodel.SetAccountId(b.AccountId)
	return svcmodel
}
