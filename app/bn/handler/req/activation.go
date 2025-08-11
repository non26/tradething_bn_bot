package req

import "tradethingbot/app/bn/process/domain"

type ActivationRequest struct {
	BotId      string `json:"botId"`
	BotOrderId string `json:"botOrderId"`
}

type ActivationRequestList []ActivationRequest

func (r *ActivationRequestList) ToDomain() []domain.Activation {
	domains := make([]domain.Activation, len(*r))
	for i, r := range *r {
		domains[i] = domain.Activation{
			BotId:      r.BotId,
			BotOrderId: r.BotOrderId,
		}
	}
	return domains
}
