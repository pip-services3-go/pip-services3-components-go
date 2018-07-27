package config

type IConfigReader interface {
	ReadConfig(correlationId string, parameters *ConfigParams) (*ConfigParams, error)
}
