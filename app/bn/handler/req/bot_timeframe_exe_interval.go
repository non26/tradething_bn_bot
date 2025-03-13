package req

import (
	"strconv"
	"strings"

	"tradethingbot/app/bn/process/domain"
)

type BotTimeframeExeIntervalHandlerRequest struct {
	BotId        string  `json:"bot_id"`
	BotOrderID   string  `json:"bot_order_id"` // client id
	Symbol       string  `json:"symbol"`
	PositionSide string  `json:"position_side"`
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
	AmountB      float64 `json:"amount_b"`
}

func (b *BotTimeframeExeIntervalHandlerRequest) Validate() error {
	return nil
}

func (b *BotTimeframeExeIntervalHandlerRequest) Transform() error {
	b.StartDate = transformToRFC3339(b.StartDate)
	b.EndDate = transformToRFC3339(b.EndDate)
	return nil
}

func transformToRFC3339(_time string) string {
	date_time := strings.Split(_time, " ")
	date := date_time[0]
	time := date_time[1]
	date_time_utc := date + "T" + time + "+07:00"
	return date_time_utc
}

func (b *BotTimeframeExeIntervalHandlerRequest) ToBotServiceRequest() *domain.BotTimeframeExeIntervalRequest {
	svcmodel := &domain.BotTimeframeExeIntervalRequest{}
	svcmodel.SetBotId(b.BotId)
	svcmodel.SetBotOrderID(b.BotOrderID)
	svcmodel.SetSymbol(b.Symbol)
	svcmodel.SetPositionSide(b.PositionSide)
	svcmodel.SetStartDate(b.StartDate)
	svcmodel.SetEndDate(b.EndDate)
	svcmodel.SetAmountB(strconv.FormatFloat(b.AmountB, 'f', -1, 64))
	return svcmodel
}
