package position

import (
	"context"
	"tradethingbot/app/bn/infrastructure/adaptor/req"

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
}

func (p *Position) ToPlacePositionModel() *req.PlacePosition {
	return &req.PlacePosition{
		PositionSide:  p.PositionSide,
		Side:          p.Side,
		EntryQuantity: p.AmountB,
		Symbol:        p.Symbol,
		ClientOrderId: p.ClientId,
	}
}

func (p *Position) ToBnFtBotOnRunTable() *dynamodbmodel.BnFtBotOnRun {
	return &dynamodbmodel.BnFtBotOnRun{
		BotID:        p.BotID,
		Symbol:       p.Symbol,
		PositionSide: p.PositionSide,
		AmountB:      p.AmountB,
		IsActive:     p.IsActive,
		BotOrderID:   p.ClientId,
	}
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
