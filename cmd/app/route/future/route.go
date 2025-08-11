package route

import (
	"tradethingbot/app/bn/handler"
	"tradethingbot/app/bn/infrastructure"
	adaptor "tradethingbot/app/bn/infrastructure/adaptor/trade"
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

	adaptorTradeOrder := adaptor.NewOrderAdaptor(
		config.BinanceFutureUrl.BaseUrl,
		config.BinanceFutureUrl.SingleOrder,
	)

	longPosition := position.NewLongPosition(
		historyTable,
		botTable,
		botOnRunTable,
		adaptorTradeOrder,
	)

	shortPosition := position.NewShortPosition(
		historyTable,
		botTable,
		botOnRunTable,
		adaptorTradeOrder,
	)

	positionBuilder := infrastructure.NewFuturePosition(
		longPosition,
		shortPosition,
	)

	trade := infrastructure.NewTrade(
		positionBuilder,
		adaptorTradeOrder,
	)

	lookUp := infrastructure.NewBotLookUp(
		botTable,
		historyTable,
		botOnRunTable,
	)

	botOnRunStore := infrastructure.NewBotOnRunStore(
		botOnRunTable,
	)

	process := process.NewBotService(
		trade,
		lookUp,
		botOnRunStore,
	)

	handler_timeframe_exe_interval := handler.NewBotTimeframeExeIntervalHandler(
		process,
	)
	path1 := "/timeframe-exe-interval"
	app.POST(path1, handler_timeframe_exe_interval.HandleBot)
	app.POST(path1+"/set", handler_timeframe_exe_interval.HandlerSetBotTimeframeExeInterval)
	app.POST(path1+"/get", handler_timeframe_exe_interval.HandlerGetBotTimeframeExeInterval)

	invalidateBotHandler := handler.NewInvalidateBotHandler(
		process,
	)
	app.POST("/invalidate-bot", invalidateBotHandler.Handle)
}
