package route

import (
	"tradethingbot/app/bn/handler"
	handlerbottimeframeexeinterval "tradethingbot/app/bn/handler/bot_timeframe_exe_interval"
	adaptor "tradethingbot/app/bn/infrastructure/adaptor/trade"
	infrastructure "tradethingbot/app/bn/infrastructure/builder"
	infralookup "tradethingbot/app/bn/infrastructure/lookup"
	"tradethingbot/app/bn/infrastructure/position"
	infrastore "tradethingbot/app/bn/infrastructure/store"
	infratrade "tradethingbot/app/bn/infrastructure/trade"
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
	botRegistor bndynamodb.IBnFtBotRegistorRepository,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
) {

	botOnRunStore := infrastore.NewBotOnRunStore(
		botOnRunTable,
	)

	botRegistorStore := infrastore.NewBotRegistorStore(
		botRegistor,
	)

	botStore := infrastore.NewBotStore(
		botTable,
	)

	historyStore := infrastore.NewBnFutureHistoryStore(
		historyTable,
	)

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

	trade := infratrade.NewTrade(
		positionBuilder,
		adaptorTradeOrder,
	)

	lookUp := infralookup.NewBotLookUp(
		botStore,
		historyStore,
		botOnRunStore,
		botRegistorStore,
	)

	process := process.NewBotService(
		trade,
		lookUp,
		botOnRunStore,
		botRegistorStore,
		config.BOTId,
	)

	handler_timeframe_exe_interval := handlerbottimeframeexeinterval.NewBotTimeframeExeIntervalHandler(
		process,
	)
	pathTimeframeExeInterval := "/timeframe-exe-interval"
	app.POST(pathTimeframeExeInterval, handler_timeframe_exe_interval.HandleBot)
	app.POST(pathTimeframeExeInterval+"/set", handler_timeframe_exe_interval.HandlerSetBotTimeframeExeInterval)
	app.GET(pathTimeframeExeInterval+"/get-all", handler_timeframe_exe_interval.HandlerGetBotTimeframeExeInterval)

	invalidateBotHandler := handler.NewInvalidateBotHandler(
		process,
	)
	app.POST("/invalidate-bot", invalidateBotHandler.Handle)

	activateHandler := handler.NewActivateHandler(
		process,
	)
	app.POST("/activate-bot", activateHandler.HandleActivate)

	deactivateHandler := handler.NewDeactivateHandler(
		process,
	)
	app.POST("/deactivate-bot", deactivateHandler.HandleDeactivate)
}
