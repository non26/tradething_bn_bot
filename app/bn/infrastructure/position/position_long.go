package position

import (
	"context"
	adaptor "tradethingbot/app/bn/infrastructure/adaptor/order"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type longPosition struct {
	historyTable  bndynamodb.IBnFtHistoryRepository
	botTable      bndynamodb.IBnFtBotRepository
	botOnRunTable bndynamodb.IBnFtBotOnRunRepository
	adaptor       adaptor.IOrderAdaptor
}

func NewLongPosition(
	historyTable bndynamodb.IBnFtHistoryRepository,
	botTable bndynamodb.IBnFtBotRepository,
	botOnRunTable bndynamodb.IBnFtBotOnRunRepository,
	adaptor adaptor.IOrderAdaptor,
) IPosition {
	return &longPosition{
		historyTable:  historyTable,
		botTable:      botTable,
		botOnRunTable: botOnRunTable,
		adaptor:       adaptor,
	}
}

func (l *longPosition) Buy(ctx context.Context, position *Position) error {

	_, err := l.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	err = l.botOnRunTable.Upsert(ctx, position.ToBnFtBotOnRunTable())
	if err != nil {
		return err
	}

	return nil
}

func (l *longPosition) Sell(ctx context.Context, position *Position) error {

	_, err := l.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	return nil
}

func (l *longPosition) Invalidate(ctx context.Context, position *Position) error {

	_, err := l.adaptor.NewOrder(ctx, position.ToPlacePositionModel())
	if err != nil {
		return err
	}

	err = l.botOnRunTable.Delete(ctx, position.ToBnFtBotOnRunTable())
	if err != nil {
		return err
	}

	err = l.historyTable.Insert(ctx, position.ToBnFtHistoryTable())
	if err != nil {
		return err
	}

	return nil
}
