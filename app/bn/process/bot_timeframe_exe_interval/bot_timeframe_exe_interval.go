package bottimeframeexeinterval

import (
	"context"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/process/domain"
)

type IBotTimeframeExeInterval interface {
	Execute(ctx context.Context, req *domain.BotTimeframeExeIntervalRequest) (*res.BotTimeframeExeIntervalResponse, error)
	Set(ctx context.Context, req *domain.BotTimeframeExeIntervalRequest) (*res.BotTimeframeExeIntervalDetailResponse, error)
	Get(ctx context.Context) ([]res.BotTimeframeExeIntervalDetailResponse, error)
}

type botTimeframeExeInterval struct {
	trade            infrastructure.ITrade
	lookUp           infrastructure.IBotLookUp
	storeBotOnRun    infrastructure.IBotOnRunStore
	storeBotRegistor infrastructure.IBotRegistorStore
}

func NewBotTimeframeExeInterval(
	trade infrastructure.ITrade,
	lookUp infrastructure.IBotLookUp,
	storeBotOnRun infrastructure.IBotOnRunStore,
	storeBotRegistor infrastructure.IBotRegistorStore,
) IBotTimeframeExeInterval {
	return &botTimeframeExeInterval{
		trade:            trade,
		lookUp:           lookUp,
		storeBotOnRun:    storeBotOnRun,
		storeBotRegistor: storeBotRegistor,
	}
}
