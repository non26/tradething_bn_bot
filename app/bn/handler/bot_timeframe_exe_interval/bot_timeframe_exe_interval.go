package handler

import "tradethingbot/app/bn/process"

type botTimeframeExeIntervalHandler struct {
	process   process.IBotService
	delayTime int
}

func NewBotTimeframeExeIntervalHandler(
	process process.IBotService,
	delayTime int,
) *botTimeframeExeIntervalHandler {
	return &botTimeframeExeIntervalHandler{
		process:   process,
		delayTime: delayTime,
	}
}
