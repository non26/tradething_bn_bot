package process

import (
	"context"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/infrastructure"
	bottimeframeexeinterval "tradethingbot/app/bn/process/bot_timeframe_exe_interval"
	"tradethingbot/app/bn/process/domain"
)

type IBotService interface {
	InvalidateBot(ctx context.Context, req *domain.InvalidateBot) (*res.InvalidateBotHandlerResponse, error)
	ActivateBot(ctx context.Context, req *domain.Activation) *res.ActivationResponse
	DeactivateBot(ctx context.Context, req *domain.Activation) *res.ActivationResponse
	DelayBot(ctx context.Context, delayTime int)
	GetBotTimeframeExeInterval() bottimeframeexeinterval.IBotTimeframeExeInterval
}

type botService struct {
	trade                   infrastructure.ITrade
	lookUp                  infrastructure.IBotLookUp
	storeBotOnRun           infrastructure.IBotOnRunStore
	storeBotRegistor        infrastructure.IBotRegistorStore
	botTimeframeExeInterval bottimeframeexeinterval.IBotTimeframeExeInterval
}

func NewBotService(
	trade infrastructure.ITrade,
	lookUp infrastructure.IBotLookUp,
	storeBotOnRun infrastructure.IBotOnRunStore,
	storeBotRegistor infrastructure.IBotRegistorStore,
) IBotService {
	botTimeframeExeInterval := bottimeframeexeinterval.NewBotTimeframeExeInterval(
		trade,
		lookUp,
		storeBotOnRun,
		storeBotRegistor,
	)

	return &botService{
		trade:                   trade,
		lookUp:                  lookUp,
		storeBotOnRun:           storeBotOnRun,
		storeBotRegistor:        storeBotRegistor,
		botTimeframeExeInterval: botTimeframeExeInterval,
	}
}

func (b *botService) GetBotTimeframeExeInterval() bottimeframeexeinterval.IBotTimeframeExeInterval {
	return b.botTimeframeExeInterval
}
