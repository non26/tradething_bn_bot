package adaptor

import (
	"context"
	"net/http"
	req "tradethingbot/app/bn/infrastructure/adaptor/req"
	res "tradethingbot/app/bn/infrastructure/adaptor/res"

	bncaller "github.com/non26/tradepkg/pkg/bn/bn_caller"
	bnrequest "github.com/non26/tradepkg/pkg/bn/bn_request"
	bnresponse "github.com/non26/tradepkg/pkg/bn/bn_response"
)

func (b *binanceFutureAdaptorService) PlaceOrder(
	ctx context.Context,
	request *req.PlacePosition,
) (*res.PlacePositionData, error) {

	c := bncaller.NewCallBinance(
		bnrequest.NewBinanceServiceHttpRequest[req.PlacePosition](),
		bnresponse.NewBinanceServiceHttpResponse[res.PlacePositionData](),
		b.httpttransport,
		b.httpclient,
	)

	res, err := c.CallBinance(
		req.NewPlaceSignleOrderBinanceServiceRequest(request),
		b.binanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1,
		b.binanceFutureUrl.SingleOrder,
		http.MethodPost,
		b.secretkey,
		b.apikey,
		b.binanceFutureServiceName,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
