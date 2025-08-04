package adaptor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	req "tradethingbot/app/bn/infrastructure/adaptor/trade/req"
	res "tradethingbot/app/bn/infrastructure/adaptor/trade/res"
)

// type NewOrderRequest struct {
// 	AccountId        string `json:"accountId" binding:"required"`
// 	PositionSide     string `json:"positionSide" binding:"required"`
// 	Side             string `json:"side" binding:"required"`
// 	Quantity         string `json:"quantity" binding:"required"`
// 	Symbol           string `json:"symbol" binding:"required"`
// 	NewClientOrderId string `json:"newClientOrderId" binding:"required"`
// 	Type             string `json:"type" binding:"required"`
// }

// type NewOrderResponse struct {
// 	// for error
// 	Code    *int    `json:"code"`
// 	Message *string `json:"msg"`
// 	// for success
// 	ClientOrderID           *string `json:"clientOrderId"`
// 	CumQty                  *string `json:"cumQty"`
// 	CumQuote                *string `json:"cumQuote"`
// 	ExecutedQty             *string `json:"executedQty"`
// 	OrderID                 *int    `json:"orderId"`
// 	AvgPrice                *string `json:"avgPrice"`
// 	OrigQty                 *string `json:"origQty"`
// 	Price                   *string `json:"price"`
// 	ReduceOnly              *bool   `json:"reduceOnly"`
// 	Side                    *string `json:"side"`
// 	PositionSide            *string `json:"positionSide"`
// 	Status                  *string `json:"status"`
// 	StopPrice               *string `json:"stopPrice"`
// 	ClosePosition           *bool   `json:"closePosition"`
// 	Symbol                  *string `json:"symbol"`
// 	TimeInForce             *string `json:"timeInForce"`
// 	Type                    *string `json:"type"`
// 	OrigType                *string `json:"origType"`
// 	ActivatePrice           *string `json:"activatePrice"`
// 	PriceRate               *string `json:"priceRate"`
// 	UpdateTime              *int64  `json:"updateTime"`
// 	WorkingType             *string `json:"workingType"`
// 	PriceProtect            *bool   `json:"priceProtect"`
// 	PriceMatch              *string `json:"priceMatch"`
// 	SelfTradePreventionMode *string `json:"selfTradePreventionMode"`
// 	GoodTillDate            *int64  `json:"goodTillDate"`
// }

func (o *orderAdaptor) NewOrder(ctx context.Context, request *req.NewOrderRequest) (*res.NewOrderResponse, error) {
	newClient := http.Client{
		Timeout: 10 * time.Second,
	}

	baseUrl := o.BaseUrl
	endpoint := o.NewOrderEndpoint
	method := http.MethodPost
	fullUrl := fmt.Sprintf("%s%s", baseUrl, endpoint)
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, fullUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := newClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var newOrderResponse res.NewOrderResponse
	err = json.Unmarshal(body, &newOrderResponse)
	if err != nil {
		return nil, err
	}

	if newOrderResponse.Code != nil && newOrderResponse.Message != nil {
		return nil, fmt.Errorf("code: %d, message: %s", *newOrderResponse.Code, *newOrderResponse.Message)
	}

	return &newOrderResponse, nil
}
