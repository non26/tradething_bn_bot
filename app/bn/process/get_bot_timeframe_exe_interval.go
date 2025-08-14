package process

import (
	"context"
	"tradethingbot/app/bn/handler/res"
)

func (b *botService) GetBotTimeframeExeInterval(ctx context.Context) ([]res.BotTimeframeExeIntervalDetailResponse, error) {
	positions, err := b.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]res.BotTimeframeExeIntervalDetailResponse, 0)
	for _, position := range positions {
		if position == nil {
			continue
		}
		detail := res.BotTimeframeExeIntervalDetailResponse{
			BotId:        position.BotID,
			BotOrderID:   position.ClientId,
			Symbol:       position.Symbol,
			PositionSide: position.PositionSide,
			AmountB:      position.AmountB,
			IsActive:     position.IsActive,
			AccountId:    position.AccountId,
		}
		result = append(result, detail)
	}
	return result, nil
}
