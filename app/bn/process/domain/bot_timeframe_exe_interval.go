package domain

import (
	"tradethingbot/app/bn/infrastructure/position"

	bnconstant "github.com/non26/tradepkg/pkg/bn/bn_constant"
	"github.com/non26/tradepkg/pkg/bn/utils"
)

type BotTimeframeExeIntervalRequest struct {
	botId        string
	botOrderID   string
	symbol       string
	positionSide string
	timeframe    string
	interval     string
	amountB      string
	// startDate    time.Time
	// endDate      time.Time
	accountId string
}

func (b *BotTimeframeExeIntervalRequest) SetBotId(botId string) {
	b.botId = botId
}

func (b *BotTimeframeExeIntervalRequest) SetBotOrderID(botOrderID string) {
	b.botOrderID = botOrderID
}

func (b *BotTimeframeExeIntervalRequest) SetSymbol(symbol string) {
	b.symbol = symbol
}

func (b *BotTimeframeExeIntervalRequest) SetPositionSide(positionSide string) {
	b.positionSide = positionSide
}

// func (b *BotTimeframeExeIntervalRequest) SetTimeframe(timeframe string) {
// 	b.timeframe = timeframe
// }

// func (b *BotTimeframeExeIntervalRequest) SetInterval(interval string) {
// 	b.interval = interval
// }

func (b *BotTimeframeExeIntervalRequest) SetAmountB(amountB string) {
	b.amountB = amountB
}

func (b *BotTimeframeExeIntervalRequest) SetAccountId(accountId string) {
	b.accountId = accountId
}

// func (b *BotTimeframeExeIntervalRequest) SetStartDate(startDate string) error {
// 	parsedTime, err := b.parseRFC3339ToUTC(startDate)
// 	if err != nil {
// 		return err
// 	}
// 	b.startDate = parsedTime
// 	return nil
// }

// func (b *BotTimeframeExeIntervalRequest) SetEndDate(endDate string) error {
// 	parsedTime, err := b.parseRFC3339ToUTC(endDate)
// 	if err != nil {
// 		return err
// 	}
// 	b.endDate = parsedTime
// 	return nil
// }

// func (b *BotTimeframeExeIntervalRequest) parseRFC3339ToUTC(_time string) (time.Time, error) {
// 	parsedTime, err := time.Parse(time.RFC3339, _time)
// 	if err != nil {
// 		return time.Time{}, err
// 	}
// 	return parsedTime.UTC(), nil
// }

// func (b *BotTimeframeExeIntervalRequest) IsPresentInTimeframe() bool {
// 	presentTime := time.Now().UTC()
// 	if b.startDate.Unix() <= presentTime.Unix() && presentTime.Unix() <= b.endDate.Unix() {
// 		return true
// 	}
// 	return false
// }

// func (b *BotTimeframeExeIntervalRequest) GetStartDate() time.Time {
// 	return b.startDate
// }

// func (b *BotTimeframeExeIntervalRequest) GetEndDate() time.Time {
// 	return b.endDate
// }

func (b *BotTimeframeExeIntervalRequest) GetBotId() string {
	return b.botId
}

func (b *BotTimeframeExeIntervalRequest) GetBotOrderID() string {
	return b.botOrderID
}

func (b *BotTimeframeExeIntervalRequest) GetSymbol() string {
	return b.symbol
}

func (b *BotTimeframeExeIntervalRequest) GetPositionSide() string {
	return b.positionSide
}

// func (b *BotTimeframeExeIntervalRequest) GetTimeframe() string {
// 	return b.timeframe
// }

// func (b *BotTimeframeExeIntervalRequest) GetInterval() string {
// 	return b.interval
// }

func (b *BotTimeframeExeIntervalRequest) GetAmountB() string {
	return b.amountB
}

func (b *BotTimeframeExeIntervalRequest) GetAccountId() string {
	return b.accountId
}

func (b *BotTimeframeExeIntervalRequest) GetOpenSide() string {
	if utils.IsLongPosition(b.positionSide) {
		return bnconstant.BUY
	}
	return bnconstant.SELL
}

func (b *BotTimeframeExeIntervalRequest) GetCloseSide() string {
	if utils.IsLongPosition(b.positionSide) {
		return bnconstant.SELL
	}
	return bnconstant.BUY
}

func (b *BotTimeframeExeIntervalRequest) ToPosition() *position.Position {
	return &position.Position{
		BotID:        b.botId,
		Symbol:       b.symbol,
		PositionSide: b.positionSide,
		AmountB:      b.amountB,
		ClientId:     b.botOrderID,
		IsActive:     true,
		AccountId:    b.accountId,
	}
}

func (b *BotTimeframeExeIntervalRequest) ToOpenPosition() *position.Position {
	side := b.GetOpenSide()
	return &position.Position{
		BotID:        b.botId,
		Symbol:       b.symbol,
		PositionSide: b.positionSide,
		AmountB:      b.amountB,
		ClientId:     b.botOrderID,
		Side:         side,
		IsActive:     true,
		AccountId:    b.accountId,
	}
}

func (b *BotTimeframeExeIntervalRequest) ToClosePosition() *position.Position {
	side := b.GetCloseSide()
	return &position.Position{
		BotID:        b.botId,
		Symbol:       b.symbol,
		PositionSide: b.positionSide,
		AmountB:      b.amountB,
		ClientId:     b.botOrderID,
		Side:         side,
		IsActive:     true,
		AccountId:    b.accountId,
	}
}

// func (b *BotTimeframeExeIntervalRequest) ToBnFtOpeningPosition() *dynamodbmodel.BnFtOpeningPosition {
// 	m := dynamodbmodel.BnFtOpeningPosition{
// 		Symbol:       b.symbol,
// 		PositionSide: b.positionSide,
// 		ClientId:     b.botOrderID,
// 		AmountB:      strconv.FormatFloat(b.amountQ, 'f', -1, 64),
// 	}
// 	return &m
// }

// func (b *BotTimeframeExeIntervalRequest) ToBnFtPlaceSingleOrderServiceRequest(side string, orderType string) *bnsvcreq.PlacePosition {
// 	m := bnsvcreq.PlacePosition{
// 		Symbol:        b.symbol,
// 		PositionSide:  b.positionSide,
// 		ClientOrderId: b.botOrderID,
// 		EntryQuantity: strconv.FormatFloat(b.amountQ, 'f', -1, 64),
// 		Side:          side,
// 		Type:          orderType,
// 	}
// 	return &m
// }

// func (b *BotTimeframeExeIntervalRequest) ToBnFtDeleteBotOnRun() *dynamodbmodel.BnFtBotOnRun {
// 	m := dynamodbmodel.BnFtBotOnRun{
// 		BotID:      b.botId,
// 		BotOrderID: b.botOrderID,
// 	}
// 	return &m
// }

// func (b *BotTimeframeExeIntervalRequest) ToBnFtHistory() *dynamodbmodel.BnFtHistory {
// 	m := dynamodbmodel.BnFtHistory{
// 		ClientId:     b.botOrderID,
// 		Symbol:       b.symbol,
// 		PositionSide: b.positionSide,
// 	}
// 	return &m
// }

// func (b *BotTimeframeExeIntervalRequest) ToBnFtBotOnRun() *dynamodbmodel.BnFtBotOnRun {
// 	m := dynamodbmodel.BnFtBotOnRun{
// 		BotID:        b.botId,
// 		BotOrderID:   b.botOrderID,
// 		Symbol:       b.symbol,
// 		PositionSide: b.positionSide,
// 		AmountB:      strconv.FormatFloat(b.amountQ, 'f', -1, 64),
// 		IsActive:     true,
// 	}
// 	return &m
// }
