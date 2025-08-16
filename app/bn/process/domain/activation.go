package domain

import (
	"tradethingbot/app/bn/infrastructure"
)

type Activation struct {
	BotId      string
	BotOrderId string
}

func (a *Activation) ToPosition() *infrastructure.Position {
	return &infrastructure.Position{
		BotID:    a.BotId,
		ClientId: a.BotOrderId,
	}
}

func (a *Activation) ToActivatedPosition() *infrastructure.Position {
	return &infrastructure.Position{
		BotID:    a.BotId,
		ClientId: a.BotOrderId,
		IsActive: true,
	}
}
