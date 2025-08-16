package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/infrastructure/position"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type botRegistorStore struct {
	botRegistorTable bndynamodb.IBnFtBotRegistorRepository
}

func NewBotRegistorStore(
	botRegistorTable bndynamodb.IBnFtBotRegistorRepository,
) infrastructure.IBotRegistorStore {
	return &botRegistorStore{
		botRegistorTable: botRegistorTable,
	}
}

func (s *botRegistorStore) Upsert(ctx context.Context, _position *position.Position) error {
	return s.botRegistorTable.Upsert(ctx, _position.ToBnFtBotRegistorTable())
}

func (s *botRegistorStore) Get(ctx context.Context, _position *position.Position) (*position.Position, error) {
	botRegistor, err := s.botRegistorTable.Get(ctx, _position.BotID, _position.ClientId)
	if err != nil {
		return nil, err
	}
	res := &position.Position{
		BotID:        botRegistor.BotID,
		ClientId:     botRegistor.BotOrderID,
		PositionSide: botRegistor.PositionSide,
		AmountB:      botRegistor.AmountQ,
		Symbol:       botRegistor.Symbol,
		AccountId:    botRegistor.AccountId,
		Setting:      []byte(botRegistor.Setting),
	}
	res.IsFoundBotRegistor = botRegistor.IsFound()
	return res, nil
}

func (s *botRegistorStore) GetAll(ctx context.Context) ([]*position.Position, error) {
	botRegistors, err := s.botRegistorTable.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	positions := make([]*position.Position, len(botRegistors))
	for i, botRegistor := range botRegistors {
		positions[i] = &position.Position{
			BotID:        botRegistor.BotID,
			ClientId:     botRegistor.BotOrderID,
			PositionSide: botRegistor.PositionSide,
			AmountB:      botRegistor.AmountQ,
			Symbol:       botRegistor.Symbol,
			AccountId:    botRegistor.AccountId,
			Setting:      []byte(botRegistor.Setting),
		}
	}
	return positions, nil
}
