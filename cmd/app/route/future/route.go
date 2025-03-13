package route

import (
	"tradethingbot/app/bn/handler"
	"tradethingbot/app/bn/infrastructure"
	"tradethingbot/app/bn/infrastructure/adaptor"
	"tradethingbot/app/bn/infrastructure/position"
	"tradethingbot/app/bn/process"
	"tradethingbot/config"

	"github.com/labstack/echo/v4"
	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

func FutureRoute(
	app *echo.Echo,
	config *config.AppConfig,
	historyTable bndynamodb.IBnFtHistoryRepository,
	bnFtCryptoTable bndynamodb.IBnFtCryptoRepository,
	botTable bndynamodb.IBnFtBotRepository,
	botOnRunTable bndynamodb.IBnFtBotOnRunRepository,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
) {
	group := app.Group("/" + config.ServiceName.BinanceFuture)

	adaptor := adaptor.NewBinanceFutureAdaptorService(
		&config.BinanceFutureUrl,
		config.Secrets.BinanceApiKey,
		config.Secrets.BinanceSecretKey,
		config.ServiceName.BinanceFuture,
		httpttransport,
		httpclient,
	)

	longPosition := position.NewLongPosition(
		historyTable,
		botTable,
		botOnRunTable,
		adaptor,
	)

	shortPosition := position.NewShortPosition(
		historyTable,
		botTable,
		botOnRunTable,
		adaptor,
	)

	positionBuilder := infrastructure.NewFuturePosition(
		longPosition,
		shortPosition,
	)

	trade := infrastructure.NewTrade(
		positionBuilder,
		adaptor,
	)

	lookUp := infrastructure.NewBotLookUp(
		botTable,
		historyTable,
		botOnRunTable,
	)

	process := process.NewBotService(
		botTable,
		botOnRunTable,
		historyTable,
		bnFtCryptoTable,
		trade,
		lookUp,
	)

	handler_timeframe_exe_interval := handler.NewBotTimeframeExeIntervalHandler(
		process,
	)
	group.POST("/timeframe-exe-interval", handler_timeframe_exe_interval.Handle)
}
