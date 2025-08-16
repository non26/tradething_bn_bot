package req

type BotRequestAspect struct {
	ActivateBotRequest   ActivateBotRequest
	DeactivateBotRequest DeactivateBotRequest
	InvalidateBotRequest InvalidateBotRequest
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
