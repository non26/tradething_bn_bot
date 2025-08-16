package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/infrastructure/position"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type botOnRunStore struct {
	botOnRunTable bndynamodb.IBnFtBotOnRunRepository
}

func NewBotOnRunStore(
	botOnRunTable bndynamodb.IBnFtBotOnRunRepository,
) infrastructure.IBotOnRunStore {
	return &botOnRunStore{
		botOnRunTable: botOnRunTable,
	}
}

func (s *botOnRunStore) Upsert(ctx context.Context, position *position.Position) error {
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
			BotID:    botOnRun.BotID,
			ClientId: botOnRun.BotOrderID,
		}
		positions = append(positions, &_position)
	}

	return positions, nil
}

func (s *botOnRunStore) Get(ctx context.Context, _position *position.Position) (*position.Position, error) {
	botOnRun, err := s.botOnRunTable.Get(ctx, _position.ToBnFtBotOnRunTable())
	if err != nil {
		return nil, err
	}

	res := &position.Position{
		BotID:    botOnRun.BotID,
		ClientId: botOnRun.BotOrderID,
	}
	res.IsFoundBotOnRunning = botOnRun.IsFound()
	return res, nil
}

func (s *botOnRunStore) Delete(ctx context.Context, position *position.Position) error {
	return s.botOnRunTable.Delete(ctx, position.ToBnFtBotOnRunTable())
}
