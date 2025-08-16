package position

import (
	"context"
	infrastructure "tradethingbot/app/bn/infrastructure"
	adaptor "tradethingbot/app/bn/infrastructure/adaptor/trade"
)

type longPosition struct {
	historyTable  infrastructure.IBnFutureHistoryStore
	botOnRunTable infrastructure.IBotOnRunStore
	adaptor       adaptor.IOrderAdaptor
}

func NewLongPosition(
	historyTable infrastructure.IBnFutureHistoryStore,
	botOnRunTable infrastructure.IBotOnRunStore,
	adaptor adaptor.IOrderAdaptor,
) infrastructure.IPosition {
	return &longPosition{
		historyTable:  historyTable,
		botOnRunTable: botOnRunTable,
		adaptor:       adaptor,
	}
}

func (l *longPosition) Buy(ctx context.Context, position *infrastructure.Position) error {

	_, err := l.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	err = l.botOnRunTable.Upsert(ctx, position)
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

	err = l.botOnRunTable.Delete(ctx, position)
	if err != nil {
		return err
	}

	err = l.historyTable.Insert(ctx, position)
	if err != nil {
		return err
	}

	return nil
}
