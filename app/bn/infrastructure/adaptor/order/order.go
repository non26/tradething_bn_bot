package adaptor

import "context"

type orderAdaptor struct {
	BaseUrl          string
	NewOrderEndpoint string
}

type IOrderAdaptor interface {
	NewOrder(ctx context.Context, request *NewOrderRequest) (*NewOrderResponse, error)
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
