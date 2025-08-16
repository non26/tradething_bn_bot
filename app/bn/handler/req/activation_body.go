package req

import "tradethingbot/app/bn/process/domain"

type ActivationRequest struct {
	BotId      string `json:"botId"`
	BotOrderId string `json:"botOrderId"`
}

func (r *ActivationRequest) ToDomain() *domain.Activation {
	return &domain.Activation{
		BotId:      r.BotId,
		BotOrderId: r.BotOrderId,
	}
}

type ActivationRequestList struct {
	ActivationRequest []ActivationRequest `json:"activationRequest"`
}
