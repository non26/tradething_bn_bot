package adaptor

import (
	"context"

	req "tradethingbot/app/bn/infrastructure/adaptor/req"
	res "tradethingbot/app/bn/infrastructure/adaptor/res"
	"tradethingbot/config"

	bnclient "github.com/non26/tradepkg/pkg/bn/bn_client"
	bntransport "github.com/non26/tradepkg/pkg/bn/bn_transport"
)

type IBinanceFutureTradeService interface {
	PlaceOrder(
		ctx context.Context,
		request *req.PlacePosition,
	) (*res.PlacePositionData, error)
}

type binanceFutureAdaptorService struct {
	binanceFutureUrl         *config.BinanceFutureUrl
	apikey                   string
	secretkey                string
	binanceFutureServiceName string
	httpttransport           bntransport.IBinanceServiceHttpTransport
	httpclient               bnclient.IBinanceSerivceHttpClient
}

func NewBinanceFutureAdaptorService(
	binanceFutureUrl *config.BinanceFutureUrl,
	apikey string,
	secretkey string,
	binanceFutureServiceName string,
	httpttransport bntransport.IBinanceServiceHttpTransport,
	httpclient bnclient.IBinanceSerivceHttpClient,
) IBinanceFutureTradeService {
	return &binanceFutureAdaptorService{
		binanceFutureUrl:         binanceFutureUrl,
		apikey:                   apikey,
		secretkey:                secretkey,
		binanceFutureServiceName: binanceFutureServiceName,
		httpttransport:           httpttransport,
		httpclient:               httpclient,
	}
}
