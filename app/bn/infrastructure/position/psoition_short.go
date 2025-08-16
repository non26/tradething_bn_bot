package position

import (
	"context"
	infrastructure "tradethingbot/app/bn/infrastructure"
	adaptor "tradethingbot/app/bn/infrastructure/adaptor/trade"
)

type shortPosition struct {
	historyTable  infrastructure.IBnFutureHistoryStore
	botOnRunTable infrastructure.IBotOnRunStore
	adaptor       adaptor.IOrderAdaptor
}

func NewShortPosition(
	historyTable infrastructure.IBnFutureHistoryStore,
	botOnRunTable infrastructure.IBotOnRunStore,
	adaptor adaptor.IOrderAdaptor,
) infrastructure.IPosition {
	return &shortPosition{
		historyTable:  historyTable,
		botOnRunTable: botOnRunTable,
		adaptor:       adaptor,
	}

}

func (s *shortPosition) Buy(ctx context.Context, position *infrastructure.Position) error {
	_, err := s.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	err = s.botOnRunTable.Upsert(ctx, position)
	if err != nil {
		return err
	}

	return nil
}

func (s *shortPosition) Sell(ctx context.Context, position *infrastructure.Position) error {
	_, err := s.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	return nil
}

func (s *shortPosition) Invalidate(ctx context.Context, position *infrastructure.Position) error {
	_, err := s.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	err = s.botOnRunTable.Delete(ctx, position)
	if err != nil {
		return err
	}

	err = s.historyTable.Insert(ctx, position)
	if err != nil {
		return err
	}

	return nil
}
