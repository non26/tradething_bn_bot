package process

import (
	"context"
	"time"
)

func (b *botService) DelayBot(ctx context.Context, delayTime int) {
	if delayTime > 0 {
		time.Sleep(time.Duration(delayTime) * time.Second)
	}
}
