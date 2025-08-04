package adaptor

import (
	"context"
	req "tradethingbot/app/bn/infrastructure/adaptor/trade/req"
	res "tradethingbot/app/bn/infrastructure/adaptor/trade/res"
)

type orderAdaptor struct {
	BaseUrl          string
	NewOrderEndpoint string
}

type IOrderAdaptor interface {
	NewOrder(ctx context.Context, request *req.NewOrderRequest) (*res.NewOrderResponse, error)
}

func NewOrderAdaptor(
	baseUrl string,
	newOrderEndpoint string,
) IOrderAdaptor {
	return &orderAdaptor{
		BaseUrl:          baseUrl,
		NewOrderEndpoint: newOrderEndpoint,
	}
}
