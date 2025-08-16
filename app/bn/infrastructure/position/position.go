package position

import (
	"context"
	adaptorreq "tradethingbot/app/bn/infrastructure/adaptor/trade/req"

	dynamodbmodel "github.com/non26/tradepkg/pkg/bn/dynamodb_future/models"
)

type Position struct {
	BotID               string
	Symbol              string
	PositionSide        string
	AmountB             string
	Side                string
	ClientId            string
	IsActive            bool
	AccountId           string
	Setting             []byte
	IsFoundInHistory    bool
	IsFoundBotID        bool
	IsFoundBotOnRunning bool
	IsFoundBotRegistor  bool
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
		BotID:      p.BotID,
		BotOrderID: p.ClientId,
	}
	return fields

}

func (p *Position) ToBnFtHistoryTable() *dynamodbmodel.BnFtHistory {
	return &dynamodbmodel.BnFtHistory{
		ClientId:     p.ClientId,
		Symbol:       p.Symbol,
		PositionSide: p.PositionSide,
	}
}

func (p *Position) ToBnFtBotRegistorTable() *dynamodbmodel.BnFtBotRegistor {
	return &dynamodbmodel.BnFtBotRegistor{
		BotID:        p.BotID,
		BotOrderID:   p.ClientId,
		PositionSide: p.PositionSide,
		AmountQ:      p.AmountB,
		Symbol:       p.Symbol,
		AccountId:    p.AccountId,
		Setting:      string(p.Setting),
	}
}

type IPosition interface {
	Buy(ctx context.Context, position *Position) error
	Sell(ctx context.Context, position *Position) error
	Invalidate(ctx context.Context, position *Position) error
}
