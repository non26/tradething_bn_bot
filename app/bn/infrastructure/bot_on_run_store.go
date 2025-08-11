package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure/position"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type botOnRunStore struct {
	botOnRunTable bndynamodb.IBnFtBotOnRunRepository
}

func NewBotOnRunStore(
	botOnRunTable bndynamodb.IBnFtBotOnRunRepository,
) IBotOnRunStore {
	return &botOnRunStore{
		botOnRunTable: botOnRunTable,
	}
}

func (s *botOnRunStore) Create(ctx context.Context, position *position.Position) error {
	return s.botOnRunTable.Upsert(ctx, position.ToBnFtBotOnRunTable())
}

func (s *botOnRunStore) GetAll(ctx context.Context) ([]*position.Position, error) {
	botOnRuns, err := s.botOnRunTable.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	positions := make([]*position.Position, len(botOnRuns))
	for _, botOnRun := range botOnRuns {
		_position := position.Position{
			BotID:        botOnRun.BotID,
			Symbol:       botOnRun.Symbol,
			PositionSide: botOnRun.PositionSide,
			AmountB:      botOnRun.AmountB,
			ClientId:     botOnRun.BotOrderID,
			IsActive:     botOnRun.IsActive,
		}
		positions = append(positions, &_position)
	}

	return positions, nil
}
