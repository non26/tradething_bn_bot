package main

import (
	"context"
	"net/http"
	"tradethingbot/cmd/app"
	route "tradethingbot/cmd/app/route/future"
	routelambda "tradethingbot/cmd/app/route/lambda"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"

	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
	bndynamodbconfig "github.com/non26/tradepkg/pkg/bn/dynamodb_config"
	bndynamodb "github.com/non26/tradepkg/pkg/bn/dynamodb_future"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	config, err := app.ReadAWSAppLog()
	if err != nil {
		panic(err.Error())
	}

	dynamodbconfig := bndynamodbconfig.NewDynamodbConfig()
	dynamodbendpoint := bndynamodbconfig.NewEndPointResolver(config.Dynamodb.Region, config.Dynamodb.Endpoint)
	dynamodbcredential := bndynamodbconfig.NewCredential(config.Dynamodb.Ak, config.Dynamodb.Sk)
	dynamodbclient := bndynamodbconfig.DynamoDB(dynamodbendpoint, dynamodbcredential, dynamodbconfig.LoadConfig()).NewPrd()

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
	routelambda.RouteLambda(app_echo, config)
	echoLambda = echoadapter.New(app_echo)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
