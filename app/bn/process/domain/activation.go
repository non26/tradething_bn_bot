package domain

import "tradethingbot/app/bn/infrastructure/position"

type Activation struct {
	BotId      string
	BotOrderId string
}

func (a *Activation) ToPosition() *position.Position {
	return &position.Position{
		BotID:    a.BotId,
		ClientId: a.BotOrderId,
	}
}

func (a *Activation) ToActivatedPosition() *position.Position {
	return &position.Position{
		BotID:    a.BotId,
		ClientId: a.BotOrderId,
		IsActive: true,
	}
}
