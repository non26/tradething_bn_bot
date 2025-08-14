package process

import (
	"context"
	"time"
	"tradethingbot/app/bn/process/domain"
)

func (b *botService) DelayBot(ctx context.Context, req *domain.BotTimeframeExeIntervalRequest) error {
	delayTime := req.GetDelayTime()
	if delayTime != nil {
		time.Sleep(time.Duration(*delayTime) * time.Second)
	}
	return nil
}
