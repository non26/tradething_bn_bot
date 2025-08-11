package domainservice

import "errors"

type ILookUpResult interface {
	ValidateBotOrderIDWith(reqBotOrderID string) error
	ValiddatePositionSideWith(reqBotPositionSide string) error
	GetAmountB() string
	IsCurrentBotActive() bool
	IsFirstTime() bool
	SetNewIsFirstTime(new bool) bool
	GetPositionSide() string
	GetSymbol() string
}

type lookUpResult struct {
	botId         string
	botOrderID    string
	positionSide  string
	amountB       string
	symbol        string
	is_active     bool
	accountId     string
	setting       string
	is_first_time bool
}

func NewLookUpResult(botId string, botOrderID string, positionSide string, amountB string, symbol string, accountId string, setting string, is_active bool) ILookUpResult {
	return &lookUpResult{
		botId:         botId,
		botOrderID:    botOrderID,
		positionSide:  positionSide,
		symbol:        symbol,
		amountB:       amountB,
		accountId:     accountId,
		setting:       setting,
		is_active:     is_active,
		is_first_time: false,
	}
}

func NewLookUpResultFirstTime(botId string, botOrderID string, positionSide string, amountB string, symbol string, accountId string, setting string, is_active bool) ILookUpResult {
	return &lookUpResult{
		botId:         botId,
		botOrderID:    botOrderID,
		positionSide:  positionSide,
		symbol:        symbol,
		amountB:       amountB,
		accountId:     accountId,
		setting:       setting,
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

func (l *lookUpResult) GetPositionSide() string {
	return l.positionSide
}

func (l *lookUpResult) GetSymbol() string {
	return l.symbol
}

func (l *lookUpResult) GetBotId() string {
	return l.botId
}

func (l *lookUpResult) GetAccountId() string {
	return l.accountId
}

func (l *lookUpResult) GetSetting() string {
	return l.setting
}
