package position

import (
	"context"
	adaptorreq "tradethingbot/app/bn/infrastructure/adaptor/trade/req"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type Position struct {
	BotID        string
	Symbol       string
	PositionSide string
	AmountB      string
	Side         string
	ClientId     string
	IsActive     bool
	AccountId    string
	Setting      []byte
}

func (p *Position) ToPlacePositionModel() *adaptorreq.NewOrderRequest {
	return &adaptorreq.NewOrderRequest{
		PositionSide:     p.PositionSide,
		Side:             p.Side,
		Quantity:         p.AmountB,
		Symbol:           p.Symbol,
		NewClientOrderId: p.ClientId,
		Type:             "MARKET",
		AccountId:        p.AccountId,
	}
}

func (p *Position) ToBnFtBotOnRunTable() *dynamodbmodel.BnFtBotOnRun {
	fields := &dynamodbmodel.BnFtBotOnRun{
		BotID:        p.BotID,
		Symbol:       p.Symbol,
		PositionSide: p.PositionSide,
		AmountB:      p.AmountB,
		IsActive:     p.IsActive,
		BotOrderID:   p.ClientId,
		AccountId:    p.AccountId,
	}
	fields.SetSetting(p.Setting)
	return fields

}

func (p *Position) ToBnFtHistoryTable() *dynamodbmodel.BnFtHistory {
	return &dynamodbmodel.BnFtHistory{
		ClientId:     p.ClientId,
		Symbol:       p.Symbol,
		PositionSide: p.PositionSide,
	}
}

type IPosition interface {
	Buy(ctx context.Context, position *Position) error
	Sell(ctx context.Context, position *Position) error
	Invalidate(ctx context.Context, position *Position) error
}
