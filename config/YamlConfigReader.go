package config

import (
	"fmt"
	"io/ioutil"

	cconfig "github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
	"gopkg.in/yaml.v2"
)

type YamlConfigReader struct {
	FileConfigReader
}

func NewEmptyYamlConfigReader() *YamlConfigReader {
	return &YamlConfigReader{
		FileConfigReader: *NewEmptyFileConfigReader(),
	}
}

func NewYamlConfigReader(path string) *YamlConfigReader {
	return &YamlConfigReader{
		FileConfigReader: *NewFileConfigReader(path),
	}
}

func (c *YamlConfigReader) ReadObject(correlationId string,
	parameters *cconfig.ConfigParams) (interface{}, error) {

	if c.Path() == "" {
		return nil, errors.NewConfigError(correlationId, "NO_PATH", "Missing config file path")
	}

	b, err := ioutil.ReadFile(c.Path())
	if err != nil {
		err = errors.NewFileError(
			correlationId,
			"READ_FAILED",
			"Failed reading configuration "+c.Path()+": "+err.Error(),
		).WithDetails("path", c.Path()).WithCause(err)
		return nil, err
	}

	data := string(b)
	data, err = c.Parameterize(data, parameters)
	if err != nil {
		return nil, err
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), m)
	if err != nil {
		return nil, err
	}

	//return convert.MapConverter.ToMap(m), nil
	return m, err
}

func (c *YamlConfigReader) ReadConfig(correlationId string,
	parameters *cconfig.ConfigParams) (result *cconfig.ConfigParams, err error) {

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	value, err := c.ReadObject(correlationId, parameters)
	if err != nil {
		return nil, err
	}

	config := cconfig.NewConfigParamsFromValue(value)
	return config, err
}

func ReadYamlObject(correlationId string, path string,
	parameters *cconfig.ConfigParams) (interface{}, error) {

	reader := NewYamlConfigReader(path)
	return reader.ReadObject(correlationId, parameters)
}

func ReadYamlConfig(correlationId string, path string,
	parameters *cconfig.ConfigParams) (*cconfig.ConfigParams, error) {

	reader := NewYamlConfigReader(path)
	return reader.ReadConfig(correlationId, parameters)
}
