package infrastructure

import (
	"context"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/infrastructure/position"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type bnFtHistoryStore struct {
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository
}

func NewBnFutureHistoryStore(
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
) infrastructure.IBnFutureHistoryStore {
	return &bnFtHistoryStore{
		bnFtHistoryTable: bnFtHistoryTable,
	}
}

func (s *bnFtHistoryStore) Get(ctx context.Context, _position *position.Position) (*position.Position, error) {
	bnFtHistory, err := s.bnFtHistoryTable.Get(ctx, _position.ClientId)
	if err != nil {
		return nil, err
	}
	res := &position.Position{
		ClientId:     bnFtHistory.ClientId,
		PositionSide: bnFtHistory.PositionSide,
	}
	res.IsFoundInHistory = bnFtHistory.IsFound()
	return res, nil
}

func (s *bnFtHistoryStore) Insert(ctx context.Context, position *position.Position) error {
	return s.bnFtHistoryTable.Insert(ctx, position.ToBnFtHistoryTable())
}
