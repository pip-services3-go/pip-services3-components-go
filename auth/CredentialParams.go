package auth

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type CredentialParams config.ConfigParams;

func NewEmptyCredentialParams() *ConfigParams {
	return data.NewEmptyStringValueMap()
}

func NewCredentialParams(values map[string]string) *CredentialParams {
	return data.NewStringValueMap(values)
}

func NewCredentialParamsFromValue(value interface{}) *CredentialParams {
	return data.NewStringValueMapFromValue(value)
}

func (c *CredentialParams) UseCredentialStore() bool {
	return c.GetAsNullableString("store_key") != nil
}

func (c *CredentialParams) StoreKey() string {
	return c.GetAsNullableString("store_key")
}

(c *CredentialParams) SetStoreKey(value string) {
	c.Put("store_key", value)
}

func (c *CredentialParams) Username() string {
	username := c.GetAsNullableString("username")
	if username == "" {
		username = c.GetAsNullableString("user")
	}
	return username
}

func (c *CredentialParams) SetUsername(value string) {
	c.Put("username", value)
}

func (c *CredentialParams) Password() string {
	password := c.GetAsNullableString("password")
	if password == "" {
		password = c.GetAsNullableString("pass")
	}
	return password
}

func (c *CredentialParams) SetPassword(value string) {
	c.Put("password", value)
}

func (c *CredentialParams) AccessId() string {
	accessId := c.GetAsNullableString("access_id")
	if accessId == "" { 
		accessId = c.GetAsNullableString("client_id")
	}
	return accessId
}

func (c *CredentialParams) SetAccessId(value string) {
	c.Put("access_id", value)
}

func (c *CredentialParams) AccessKey() string {
	accessKey := c.GetAsNullableString("access_key")
	if accessKey == "" {
		accessKey = c.GetAsNullableString("client_key")
	}
	return accessKey
}

func (c *CredentialParams) SetAccessKey(value string) {
	c.Put("access_key", value)
}

func NewCredentialParamsFromString(line string) *CredentialParams {
	return NewStringValueMapFromString(line)
}

func NewManyCredentialParamsFromConfig(config *ConfigParams) []*CredentialParams {
	result := []CredentialParams {}

	credentials := config.GetSection("credentials")

	if len(credentials) > 0 {
		for _, section := range credentials.GetSectionNames() {
			credential := credentials.GetSection(section)
			result = append(result, NewCredentialParams(credential))
		}
	} else {
		credential := config.GetSection("credential")
		if len(credential) > 0 {
			result = append(result, NewCredentialParams(credential))
		}
	}

	return result
}

func NewCredentialParamsFromConfig(config ConfigParams) *CredentialParams {
	credentials := NewManyCredentialParamsFromConfig(config)
	if len(credentials) > 0 {
		return credentials[0]
	 } else {
		 return nil
	 }
}