package config

import c "github.com/pip-services3-go/pip-services3-commons-go/config"

type IConfigReader interface {
	ReadConfig(correlationId string, parameters *c.ConfigParams) (*c.ConfigParams, error)
}
