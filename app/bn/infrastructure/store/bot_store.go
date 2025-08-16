package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/infrastructure/position"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type botStore struct {
	botTable bndynamodb.IBnFtBotRepository
}

func NewBotStore(
	botTable bndynamodb.IBnFtBotRepository,
) infrastructure.IBotStore {
	return &botStore{
		botTable: botTable,
	}
}

func (s *botStore) Get(ctx context.Context, botId string) (*position.Position, error) {
	bot, err := s.botTable.Get(ctx, botId)
	if err != nil {
		return nil, err
	}
	res := &position.Position{
		BotID: bot.BotID,
	}
	res.IsFoundBotID = bot.IsFound()
	return res, nil
}
