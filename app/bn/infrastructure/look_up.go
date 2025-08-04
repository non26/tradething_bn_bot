package infrastructure

import (
	"context"
	"errors"
	"tradethingbot/app/bn/infrastructure/position"
	domainservice "tradethingbot/app/bn/process/domain_service"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type botLookUp struct {
	botTable      bndynamodb.IBnFtBotRepository
	historyTable  bndynamodb.IBnFtHistoryRepository
	botOnRunTable bndynamodb.IBnFtBotOnRunRepository
}

func NewBotLookUp(
	botTable bndynamodb.IBnFtBotRepository,
	historyTable bndynamodb.IBnFtHistoryRepository,
	botOnRunTable bndynamodb.IBnFtBotOnRunRepository,
) IBotLookUp {
	return &botLookUp{
		botTable:      botTable,
		historyTable:  historyTable,
		botOnRunTable: botOnRunTable,
	}
}

func (b *botLookUp) LookUp(ctx context.Context, position *position.Position) (result domainservice.ILookUpResult, err error) {

	bot, err := b.botTable.Get(ctx, position.BotID)
	if err != nil {
		return nil, err
	}
	if !bot.IsFound() {
		return nil, errors.New("bot not found")
	}

	posHistory, err := b.historyTable.Get(ctx, position.ClientId)
	if err != nil {
		return nil, err
	}
	if posHistory.IsFound() {
		return nil, errors.New("bot order already closed")
	}

	current_position, err := b.botOnRunTable.Get(ctx, position.ToBnFtBotOnRunTable())
	if err != nil {
		return nil, err
	}
	if current_position.IsFound() {
		return domainservice.NewLookUpResult(current_position.BotOrderID, current_position.PositionSide, current_position.AmountB, current_position.IsActive), nil
	}

	return domainservice.NewLookUpResultFirstTime(current_position.BotOrderID, current_position.PositionSide, current_position.IsActive), nil
}
