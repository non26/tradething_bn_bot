package process

import (
	"context"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/process/domain"
)

type IBotService interface {
	InvalidateBot(ctx context.Context, req *domain.InvalidateBot) (*res.InvalidateBotHandlerResponse, error)
	BotTimeframeExeInterval(ctx context.Context, req *domain.BotTimeframeExeIntervalRequest) (*res.BotTimeframeExeIntervalResponse, error)
	SetBotTimeframeExeInterval(ctx context.Context, req *domain.BotTimeframeExeIntervalRequest) (*res.BotTimeframeExeIntervalDetailResponse, error)
	GetBotTimeframeExeInterval(ctx context.Context) ([]res.BotTimeframeExeIntervalDetailResponse, error)
	ActivateBot(ctx context.Context, req []domain.Activation) ([]res.ActivationResponse, error)
	DeactivateBot(ctx context.Context, req []domain.Activation) ([]res.ActivationResponse, error)
}

type botService struct {
	trade  infrastructure.ITrade
	lookUp infrastructure.IBotLookUp
	store  infrastructure.IBotOnRunStore
}

func NewBotService(
	trade infrastructure.ITrade,
	lookUp infrastructure.IBotLookUp,
	store infrastructure.IBotOnRunStore,
) IBotService {
	return &botService{
		trade:  trade,
		lookUp: lookUp,
		store:  store,
	}
}
