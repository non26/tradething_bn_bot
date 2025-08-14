package req

import "tradethingbot/app/bn/process/domain"

type ActivationRequest struct {
	BotId      string `json:"botId"`
	BotOrderId string `json:"botOrderId"`
}

type ActivationRequestList struct {
	ActivationRequest []ActivationRequest `json:"activationRequest"`
}

func (r *ActivationRequestList) ToDomain() []domain.Activation {
	domains := make([]domain.Activation, len(r.ActivationRequest))
	for i, request := range r.ActivationRequest {
		domains[i] = domain.Activation{
			BotId:      request.BotId,
			BotOrderId: request.BotOrderId,
		}
	}
	return domains
}
