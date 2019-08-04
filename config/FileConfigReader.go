package config

import cconfig "github.com/pip-services3-go/pip-services3-commons-go/config"

type FileConfigReader struct {
	ConfigReader
	path	string
}

func NewEmptyFileConfigReader() *FileConfigReader {
	return &FileConfigReader {
		ConfigReader: *NewConfigReader(),
	}
}

func NewFileConfigReader(path string) *FileConfigReader {
	return &FileConfigReader {
		ConfigReader: *NewConfigReader(),
		path: path,
	}
}

func (c *FileConfigReader) Configure(config *cconfig.ConfigParams) {
	c.ConfigReader.Configure(config)
	c.path = config.GetAsStringWithDefault("path", c.path)
}

func (c *FileConfigReader) Path() string {
	return c.path
}

func (c *FileConfigReader) SetPath(path string) {
	c.path = path
}