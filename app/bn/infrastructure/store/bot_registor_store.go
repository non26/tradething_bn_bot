package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure"

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

func (s *botRegistorStore) Upsert(ctx context.Context, _position *infrastructure.Position) error {
	return s.botRegistorTable.Upsert(ctx, _position.ToBnFtBotRegistorTable())
}

func (s *botRegistorStore) Get(ctx context.Context, _position *infrastructure.Position) (*infrastructure.Position, error) {
	botRegistor, err := s.botRegistorTable.Get(ctx, _position.BotID, _position.ClientId)
	if err != nil {
		return nil, err
	}
	res := &infrastructure.Position{
		BotID:        botRegistor.BotID,
		ClientId:     botRegistor.BotOrderID,
		PositionSide: botRegistor.PositionSide,
		AmountB:      botRegistor.AmountQ,
		Symbol:       botRegistor.Symbol,
		AccountId:    botRegistor.AccountId,
		Setting:      []byte(botRegistor.Setting),
		IsActive:     botRegistor.IsActive,
	}
	res.IsFoundBotRegistor = botRegistor.IsFound()
	return res, nil
}

func (s *botRegistorStore) GetAll(ctx context.Context) ([]*infrastructure.Position, error) {
	botRegistors, err := s.botRegistorTable.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	positions := make([]*infrastructure.Position, len(botRegistors))
	for i, botRegistor := range botRegistors {
		positions[i] = &infrastructure.Position{
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

func (s *botRegistorStore) Delete(ctx context.Context, _position *infrastructure.Position) error {
	return s.botRegistorTable.Delete(ctx, _position.BotID, _position.ClientId)
}
