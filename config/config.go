package config

type AppConfig struct {
	Env              string        `mapstructure:"environment" json:"environment"`
	Port             string        `mapstructure:"port" json:"port"`
	BinanceFutureUrl BinanceFuture `mapstructure:"binance-future" json:"binance-future"`
	Dynamodb         Dynamodb      `mapstructure:"dynamodb" json:"dynamodb"`
	BOTId            BOTId         `mapstructure:"bot-id" json:"bot-id"`
}

type Dynamodb struct {
	Region   string `mapstructure:"region" json:"region"`
	Ak       string `mapstructure:"ak" json:"ak"`
	Sk       string `mapstructure:"sk" json:"sk"`
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
}

type BinanceFuture struct {
	DelayTime     int    `mapstructure:"delay-time" json:"delay-time"`
	SetLeverage   string `mapstructure:"set-leverage" json:"set-leverage"`
	SingleOrder   string `mapstructure:"single-order" json:"single-order"`
	MultipleOrder string `mapstructure:"miltiple-order" json:"miltiple-order"`
	QueryOrder    string `mapstructure:"query-order" json:"query-order"`
	ExchangeInfo  string `mapstructure:"exchange-info" json:"exchange-info"`
	BaseUrl       string `mapstructure:"baseUrl" json:"baseUrl"`
}

type BOTId struct {
	TimeFrameExeIntervalId string `mapstructure:"time-frame-exe-interval-id" json:"time-frame-exe-interval-id"`
}
