package domainservice

import "errors"

type ILookUpResult interface {
	ValidateBotOrderIDWith(reqBotOrderID string) error
	ValiddatePositionSideWith(reqBotPositionSide string) error
	GetAmountB() string
	IsCurrentBotActive() bool
	IsFirstTime() bool
	SetNewIsFirstTime(new bool) bool
}

type lookUpResult struct {
	botOrderID    string
	positionSide  string
	amountB       string
	is_active     bool
	is_first_time bool
}

func NewLookUpResult(botOrderID string, positionSide string, amountB string, is_active bool) ILookUpResult {
	return &lookUpResult{
		botOrderID:    botOrderID,
		positionSide:  positionSide,
		amountB:       amountB,
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

func (l *lookUpResult) GetAmountB() string {
	return l.amountB
}

func (l *lookUpResult) SetNewIsFirstTime(new bool) bool {
	l.is_first_time = new
	return l.is_first_time
}
