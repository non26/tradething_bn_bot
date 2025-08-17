package position

import (
	"context"
	infrastructure "tradethingbot/app/bn/infrastructure"
	adaptor "tradethingbot/app/bn/infrastructure/adaptor/trade"
)

type shortPosition struct {
	historyStore     infrastructure.IBnFutureHistoryStore
	botOnRunStore    infrastructure.IBotOnRunStore
	botRegistorStore infrastructure.IBotRegistorStore
	adaptor          adaptor.IOrderAdaptor
}

func NewShortPosition(
	historyStore infrastructure.IBnFutureHistoryStore,
	botOnRunStore infrastructure.IBotOnRunStore,
	botRegistorStore infrastructure.IBotRegistorStore,
	adaptor adaptor.IOrderAdaptor,
) infrastructure.IPosition {
	return &shortPosition{
		historyStore:     historyStore,
		botOnRunStore:    botOnRunStore,
		botRegistorStore: botRegistorStore,
		adaptor:          adaptor,
	}

}

func (s *shortPosition) Buy(ctx context.Context, position *infrastructure.Position) error {
	_, err := s.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	err = s.botOnRunStore.Upsert(ctx, position)
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

	err = s.botOnRunStore.Delete(ctx, position)
	if err != nil {
		return err
	}

	err = s.botRegistorStore.Delete(ctx, position)
	if err != nil {
		return err
	}

	err = s.historyStore.Insert(ctx, position)
	if err != nil {
		return err
	}

	return nil
}
