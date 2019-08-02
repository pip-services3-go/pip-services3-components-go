package auth

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type CredentialParams struct {
	config.ConfigParams
}

func NewEmptyCredentialParams() *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewEmptyConfigParams(),
	}
}

func NewCredentialParams(values map[string]string) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParams(values),
	}
}

func NewCredentialParamsFromValue(value interface{}) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParamsFromValue(value),
	}
}

func NewCredentialParamsFromTuples(tuples ...interface{}) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParamsFromTuplesArray(tuples),
	}
}

func NewCredentialParamsFromTuplesArray(tuples []interface{}) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParamsFromTuplesArray(tuples),
	}
}

func NewCredentialParamsFromString(line string) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParamsFromString(line),
	}
}

func NewCredentialParamsFromMaps(maps ...map[string]string) *CredentialParams {
	return &CredentialParams{
		ConfigParams: *config.NewConfigParamsFromMaps(maps...),
	}
}

func NewManyCredentialParamsFromConfig(config *config.ConfigParams) []*CredentialParams {
	result := []*CredentialParams{}

	credentials := config.GetSection("credentials")

	if credentials.Len() > 0 {
		for _, section := range credentials.GetSectionNames() {
			credential := credentials.GetSection(section)
			result = append(result, NewCredentialParams(credential.Value()))
		}
	} else {
		credential := config.GetSection("credential")
		if credential.Len() > 0 {
			result = append(result, NewCredentialParams(credential.Value()))
		}
	}

	return result
}

func NewCredentialParamsFromConfig(config *config.ConfigParams) *CredentialParams {
	credentials := NewManyCredentialParamsFromConfig(config)
	if len(credentials) > 0 {
		return credentials[0]
	}
	return nil
}

func (c *CredentialParams) UseCredentialStore() bool {
	return c.GetAsString("store_key") != ""
}

func (c *CredentialParams) StoreKey() string {
	return c.GetAsString("store_key")
}

func (c *CredentialParams) SetStoreKey(value string) {
	c.Put("store_key", value)
}

func (c *CredentialParams) Username() string {
	username := c.GetAsString("username")
	if username == "" {
		username = c.GetAsString("user")
	}
	return username
}

func (c *CredentialParams) SetUsername(value string) {
	c.Put("username", value)
}

func (c *CredentialParams) Password() string {
	password := c.GetAsString("password")
	if password == "" {
		password = c.GetAsString("pass")
	}
	return password
}

func (c *CredentialParams) SetPassword(value string) {
	c.Put("password", value)
}

func (c *CredentialParams) AccessId() string {
	accessId := c.GetAsString("access_id")
	if accessId == "" {
		accessId = c.GetAsString("client_id")
	}
	return accessId
}

func (c *CredentialParams) SetAccessId(value string) {
	c.Put("access_id", value)
}

func (c *CredentialParams) AccessKey() string {
	accessKey := c.GetAsString("access_key")
	if accessKey == "" {
		accessKey = c.GetAsString("client_key")
	}
	return accessKey
}

func (c *CredentialParams) SetAccessKey(value string) {
	c.Put("access_key", value)
}
