package position

import (
	"context"
	adaptor "tradethingbot/app/bn/infrastructure/adaptor/trade"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type shortPosition struct {
	historyTable  bndynamodb.IBnFtHistoryRepository
	botTable      bndynamodb.IBnFtBotRepository
	botOnRunTable bndynamodb.IBnFtBotOnRunRepository
	adaptor       adaptor.IOrderAdaptor
}

func NewShortPosition(
	historyTable bndynamodb.IBnFtHistoryRepository,
	botTable bndynamodb.IBnFtBotRepository,
	botOnRunTable bndynamodb.IBnFtBotOnRunRepository,
	adaptor adaptor.IOrderAdaptor,
) IPosition {
	return &shortPosition{
		historyTable:  historyTable,
		botTable:      botTable,
		botOnRunTable: botOnRunTable,
		adaptor:       adaptor,
	}

}

func (s *shortPosition) Buy(ctx context.Context, position *Position) error {
	_, err := s.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	err = s.botOnRunTable.Upsert(ctx, position.ToBnFtBotOnRunTable())
	if err != nil {
		return err
	}

	return nil
}

func (s *shortPosition) Sell(ctx context.Context, position *Position) error {
	_, err := s.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	return nil
}

func (s *shortPosition) Invalidate(ctx context.Context, position *Position) error {
	_, err := s.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	err = s.botOnRunTable.Delete(ctx, position.ToBnFtBotOnRunTable())
	if err != nil {
		return err
	}

	err = s.historyTable.Insert(ctx, position.ToBnFtHistoryTable())
	if err != nil {
		return err
	}

	return nil
}
