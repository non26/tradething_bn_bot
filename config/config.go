package config

type AppConfig struct {
	Env              string           `mapstructure:"environment" json:"environment"`
	Port             string           `mapstructure:"port" json:"port"`
	ServiceName      ServiceName      `mapstructure:"service-name" json:"service-name"`
	Secrets          Secrets          `mapstructure:"secrets" json:"secrets"`
	BinanceFutureUrl BinanceFutureUrl `mapstructure:"binance-future-url" json:"binance-future-url"`
	Dynamodb         Dynamodb         `mapstructure:"dynamodb" json:"dynamodb"`
}

type Dynamodb struct {
	Region   string `mapstructure:"region" json:"region"`
	Ak       string `mapstructure:"ak" json:"ak"`
	Sk       string `mapstructure:"sk" json:"sk"`
	Endpoint string `mapstructure:"endpoint" json:"endpoint"`
}

type Secrets struct {
	BinanceApiKey    string `mapstructure:"binance-apiKey" json:"binance-apiKey"`
	BinanceSecretKey string `mapstructure:"binance-secretKey" json:"binance-secretKey"`
}

type ServiceName struct {
	BinanceFuture string `mapstructure:"binance-future" json:"binance-future"`
}

type BinanceFutureBaseUrl struct {
	BianceUrl1 string `mapstructure:"binance1" json:"binance1"`
}

type BinanceFutureUrl struct {
	SetLeverage          string               `mapstructure:"set-leverage" json:"set-leverage"`
	SingleOrder          string               `mapstructure:"single-order" json:"single-order"`
	MultipleOrder        string               `mapstructure:"miltiple-order" json:"miltiple-order"`
	QueryOrder           string               `mapstructure:"query-order" json:"query-order"`
	ExchangeInfo         string               `mapstructure:"exchange-info" json:"exchange-info"`
	BinanceFutureBaseUrl BinanceFutureBaseUrl `mapstructure:"binance-future-baseUrl" json:"binance-future-baseUrl"`
}
