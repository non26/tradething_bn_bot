package infrastructure

import (
	"context"
	"errors"
	"tradethingbot/app/bn/infrastructure"
	domainservice "tradethingbot/app/bn/process/domain_service"
)

type botLookUp struct {
	botTable      infrastructure.IBotStore
	historyTable  infrastructure.IBnFutureHistoryStore
	botOnRunTable infrastructure.IBotOnRunStore
	botRegistor   infrastructure.IBotRegistorStore
}

func NewBotLookUp(
	botTable infrastructure.IBotStore,
	historyTable infrastructure.IBnFutureHistoryStore,
	botOnRunTable infrastructure.IBotOnRunStore,
	botRegistor infrastructure.IBotRegistorStore,
) infrastructure.IBotLookUp {
	return &botLookUp{
		botTable:      botTable,
		historyTable:  historyTable,
		botOnRunTable: botOnRunTable,
		botRegistor:   botRegistor,
	}
}

func (b *botLookUp) LookUp(ctx context.Context, position *infrastructure.Position) (result domainservice.ILookUpResult, err error) {

	bot, err := b.botTable.Get(ctx, position.BotID)
	if err != nil {
		return nil, err
	}
	if !bot.IsFoundBotID {
		return nil, errors.New("bot not found")
	}

	positionHistory, err := b.historyTable.Get(ctx, position)
	if err != nil {
		return nil, err
	}
	if positionHistory.IsFoundInHistory {
		return nil, errors.New("bot order already closed")
	}

	botRegistor, err := b.botRegistor.Get(ctx, position)
	if err != nil {
		return nil, err
	}

	if !botRegistor.IsFoundBotRegistor {
		return domainservice.NewLookUpResultRegistor(
			false,
		), nil
	}

	botOnRun, err := b.botOnRunTable.Get(ctx, position)
	if err != nil {
		return nil, err
	}

	if botOnRun.IsFoundBotOnRunning {
		return domainservice.NewLookUpResult(
			botOnRun.BotID,
			botOnRun.ClientId,
			botRegistor.PositionSide,
			botRegistor.AmountB,
			botRegistor.Symbol,
			botRegistor.AccountId,
			string(botRegistor.Setting),
			botRegistor.IsActive,
			botRegistor.IsFoundBotRegistor,
		), nil
	}

	return domainservice.NewLookUpResultFirstTime(
		botOnRun.BotID,
		botOnRun.ClientId,
		botRegistor.PositionSide,
		botRegistor.AmountB,
		botRegistor.Symbol,
		botRegistor.AccountId,
		string(botRegistor.Setting),
		botRegistor.IsActive,
		botRegistor.IsFoundBotRegistor,
	), nil
}
