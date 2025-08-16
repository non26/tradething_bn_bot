package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure"

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

func (s *botOnRunStore) Upsert(ctx context.Context, position *infrastructure.Position) error {
	return s.botOnRunTable.Upsert(ctx, position.ToBnFtBotOnRunTable())
}

func (s *botOnRunStore) GetAll(ctx context.Context) ([]*infrastructure.Position, error) {
	botOnRuns, err := s.botOnRunTable.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	positions := make([]*infrastructure.Position, len(botOnRuns))
	for _, botOnRun := range botOnRuns {
		_position := infrastructure.Position{
			BotID:    botOnRun.BotID,
			ClientId: botOnRun.BotOrderID,
		}
		positions = append(positions, &_position)
	}

	return positions, nil
}

func (s *botOnRunStore) Get(ctx context.Context, _position *infrastructure.Position) (*infrastructure.Position, error) {
	botOnRun, err := s.botOnRunTable.Get(ctx, _position.ToBnFtBotOnRunTable())
	if err != nil {
		return nil, err
	}

	res := &infrastructure.Position{
		BotID:    botOnRun.BotID,
		ClientId: botOnRun.BotOrderID,
	}
	res.IsFoundBotOnRunning = botOnRun.IsFound()
	return res, nil
}

func (s *botOnRunStore) Delete(ctx context.Context, position *infrastructure.Position) error {
	return s.botOnRunTable.Delete(ctx, position.ToBnFtBotOnRunTable())
}
