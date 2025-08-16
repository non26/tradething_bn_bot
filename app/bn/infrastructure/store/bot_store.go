package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure"

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

func (s *botStore) Get(ctx context.Context, botId string) (*infrastructure.Position, error) {
	bot, err := s.botTable.Get(ctx, botId)
	if err != nil {
		return nil, err
	}
	res := &infrastructure.Position{
		BotID: bot.BotID,
	}
	res.IsFoundBotID = bot.IsFound()
	return res, nil
}
