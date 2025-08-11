package res

type ActivationResponse struct {
	Status     string `json:"status"`
	Message    string `json:"message,omitempty"`
	BotId      string `json:"botId"`
	BotOrderId string `json:"botOrderId"`
}

func (a *ActivationResponse) SetFailed(message string) {
	a.Status = "failed"
	a.Message = message
}

func (a *ActivationResponse) SetSuccess() {
	a.Status = "success"

}
