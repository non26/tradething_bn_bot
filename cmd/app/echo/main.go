package main

import (
	"fmt"
	"net/http"
	"tradethingbot/cmd/app"
	route "tradethingbot/cmd/app/route/future"

	"github.com/labstack/echo/v4"

	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	bndynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

func main() {

	config, err := app.ReadLog("./config")
	if err != nil {
		panic(err.Error())
	}

	dynamodbconfig := bndynamodbconfig.NewDynamodbConfig()
	dynamodbendpoint := bndynamodbconfig.NewEndPointResolver(config.Dynamodb.Region, config.Dynamodb.Endpoint)
	dynamodbcredential := bndynamodbconfig.NewCredential(config.Dynamodb.Ak, config.Dynamodb.Sk)
	dynamodbclient := bndynamodbconfig.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewLocal()

	bnFtCryptoTable := bndynamodb.NewConnectionBnFtCryptoRepository(dynamodbclient)
	bnFtHistoryTable := bndynamodb.NewConnectionBnFtHistoryRepository(dynamodbclient)
	bnFtBotTable := bndynamodb.NewConnectionBnFtBotRepository(dynamodbclient)
	bnFtBotOnRunTable := bndynamodb.NewConnectionBnFtBotOnRunRepository(dynamodbclient)
	bnFtBotRegistorTable := bndynamodb.NewConnectionBnFtBotRegistorRepository(dynamodbclient)

	httptransport := bntransport.NewBinanceTransport(&http.Transport{})
	httpclient := bnclient.NewBinanceSerivceHttpClient()

	app_echo := echo.New()
	app.HealthCheck(app_echo)
	route.FutureRoute(
		app_echo,
		config,
		bnFtHistoryTable,
		bnFtCryptoTable,
		bnFtBotTable,
		bnFtBotOnRunTable,
		bnFtBotRegistorTable,
		httptransport,
		httpclient,
	)
	port := fmt.Sprintf(":%v", config.Port)
	app_echo.Start(port)
}
