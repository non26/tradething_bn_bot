package position

import (
	"context"
	infrastructure "tradethingbot/app/bn/infrastructure"
	adaptor "tradethingbot/app/bn/infrastructure/adaptor/trade"
)

type longPosition struct {
	historyStore     infrastructure.IBnFutureHistoryStore
	botOnRunStore    infrastructure.IBotOnRunStore
	botRegistorStore infrastructure.IBotRegistorStore
	adaptor          adaptor.IOrderAdaptor
}

func NewLongPosition(
	historyStore infrastructure.IBnFutureHistoryStore,
	botOnRunStore infrastructure.IBotOnRunStore,
	botRegistorStore infrastructure.IBotRegistorStore,
	adaptor adaptor.IOrderAdaptor,
) infrastructure.IPosition {
	return &longPosition{
		historyStore:     historyStore,
		botOnRunStore:    botOnRunStore,
		botRegistorStore: botRegistorStore,
		adaptor:          adaptor,
	}
}

func (l *longPosition) Buy(ctx context.Context, position *infrastructure.Position) error {

	_, err := l.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	err = l.botOnRunStore.Upsert(ctx, position)
	if err != nil {
		return err
	}

	return nil
}

func (l *longPosition) Sell(ctx context.Context, position *infrastructure.Position) error {

	_, err := l.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	return nil
}

func (l *longPosition) Invalidate(ctx context.Context, position *infrastructure.Position) error {

	_, err := l.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	err = l.botOnRunStore.Delete(ctx, position)
	if err != nil {
		return err
	}

	err = l.botRegistorStore.Delete(ctx, position)
	if err != nil {
		return err
	}

	err = l.historyStore.Insert(ctx, position)
	if err != nil {
		return err
	}

	return nil
}
