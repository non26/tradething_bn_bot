package adaptor

import (
	"context"
	req "tradethingbot/app/bn/infrastructure/adaptor/order/req"
	res "tradethingbot/app/bn/infrastructure/adaptor/order/res"
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
