package req

import (
	"tradethingbot/app/bn/process/domain"
)

type BotTimeframeExeIntervalHandlerRequest struct {
	BotId        string `json:"botId"`
	BotOrderID   string `json:"botOrderId"` // client id
	Symbol       string `json:"symbol"`
	PositionSide string `json:"positionSide"`
	AmountB      string `json:"amountB"`
	AccountId    string `json:"accountId"`
}

func (b *BotTimeframeExeIntervalHandlerRequest) Validate() error {
	return nil
}

func (b *BotTimeframeExeIntervalHandlerRequest) Transform() error {
	// b.StartDate = transformToRFC3339(b.StartDate)
	// b.EndDate = transformToRFC3339(b.EndDate)
	return nil
}

// func transformToRFC3339(_time string) string {
// 	date_time := strings.Split(_time, " ")
// 	date := date_time[0]
// 	time := date_time[1]
// 	date_time_utc := date + "T" + time + "+07:00"
// 	return date_time_utc
// }

func (b *BotTimeframeExeIntervalHandlerRequest) ToBotServiceRequest() *domain.BotTimeframeExeIntervalRequest {
	svcmodel := &domain.BotTimeframeExeIntervalRequest{}
	svcmodel.SetBotId(b.BotId)
	svcmodel.SetBotOrderID(b.BotOrderID)
	svcmodel.SetSymbol(b.Symbol)
	svcmodel.SetPositionSide(b.PositionSide)
	svcmodel.SetAmountB(b.AmountB)
	svcmodel.SetAccountId(b.AccountId)
	return svcmodel
}
