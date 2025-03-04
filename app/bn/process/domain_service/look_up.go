package domainservice

import "errors"

type ILookUpResult interface {
	ValidateBotOrderIDWith(reqBotOrderID string) error
	ValiddatePositionSideWith(reqBotPositionSide string) error
	IsCurrentBotActive() bool
	IsFirstTime() bool
}

type lookUpResult struct {
	botOrderID    string
	positionSide  string
	is_active     bool
	is_first_time bool
}

func NewLookUpResult(botOrderID string, positionSide string, is_active bool) ILookUpResult {
	return &lookUpResult{
		botOrderID:    botOrderID,
		positionSide:  positionSide,
		is_active:     is_active,
		is_first_time: false,
	}
}

func NewLookUpResultFirstTime(botOrderID string, positionSide string, is_active bool) ILookUpResult {
	return &lookUpResult{
		botOrderID:    botOrderID,
		positionSide:  positionSide,
		is_active:     is_active,
		is_first_time: true,
	}
}
func (l *lookUpResult) ValidateBotOrderIDWith(reqBotOrderID string) error {
	if l.botOrderID != reqBotOrderID {
		return errors.New("bot order id not match")
	}
	return nil
}

func (l *lookUpResult) ValiddatePositionSideWith(reqBotPositionSide string) error {
	if l.positionSide != reqBotPositionSide {
		return errors.New("position side not match")
	}
	return nil
}

func (l *lookUpResult) IsCurrentBotActive() bool {
	return l.is_active
}

func (l *lookUpResult) IsFirstTime() bool {
	return l.is_first_time
}
