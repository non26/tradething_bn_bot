package req

type BotRequestAspect struct {
	ActivateBotRequest   ActivateBotRequest
	DeactivateBotRequest DeactivateBotRequest
	InvalidateBotRequest InvalidateBotRequest
	DelayBotRequest      DelayBotRequest
}

type ActivateBotRequest struct {
	ClientIds []string `json:"clientIds"`
}

type DeactivateBotRequest struct {
	ClientIds []string `json:"clientIds"`
}

type InvalidateBotRequest struct {
	ClientIds []string `json:"clientIds"`
}

type DelayBotRequest struct {
	DelayTime *int `json:"delayTime"`
}
