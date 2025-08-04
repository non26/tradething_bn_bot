package process

import (
	"context"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/process/domain"

	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

type IBotService interface {
	InvalidateBot(ctx context.Context, req *domain.InvalidateBot) (*res.InvalidateBotHandlerResponse, error)
	BotTimeframeExeInterval(ctx context.Context, req *domain.BotTimeframeExeIntervalRequest) (*res.BotTimeframeExeIntervalResponse, error)
}

type botService struct {
	// binanceService     bntrade.IBinanceFutureExternalService
	bnFtBotTable      bndynamodb.IBnFtBotRepository
	bnFtBotOnRunTable bndynamodb.IBnFtBotOnRunRepository
	bnFtHistoryTable  bndynamodb.IBnFtHistoryRepository
	bnFtCryptoTable   bndynamodb.IBnFtCryptoRepository
	trade             infrastructure.ITrade
	lookUp            infrastructure.IBotLookUp
}

func NewBotService(
	bnFtBotTable bndynamodb.IBnFtBotRepository,
	bnFtBotOnRunTable bndynamodb.IBnFtBotOnRunRepository,
	bnFtHistoryTable bndynamodb.IBnFtHistoryRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	trade infrastructure.ITrade,
	lookUp infrastructure.IBotLookUp,
) IBotService {
	return &botService{
		bnFtBotTable:      bnFtBotTable,
		bnFtBotOnRunTable: bnFtBotOnRunTable,
		bnFtHistoryTable:  bnFtHistoryTable,
		bnFtCryptoTable:   bnFtCryptoTable,
		trade:             trade,
		lookUp:            lookUp,
	}
}
