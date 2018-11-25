package auth

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/data"
)

//TODO::Need to resolve names of package
type CredentialParams config.ConfigParams

func NewEmptyCredentialParams() *CredentialParams {
	return &CredentialParams{
		StringValueMap: *data.NewEmptyStringValueMap(),
	}
}

func NewCredentialParams(values map[string]string) *CredentialParams {
	return &CredentialParams{
		StringValueMap: *data.NewStringValueMap(values),
	}
}

func NewCredentialParamsFromValue(value interface{}) *CredentialParams {
	return &CredentialParams{
		StringValueMap: *data.NewStringValueMapFromValue(value),
	}
}

func NewCredentialParamsManyFromConfig(config config.ConfigParams) *[]CredentialParams {
	result := []CredentialParams{}

	credentials := config.GetSection("credentials")

	if credentials.Len() > 0 {
		for _, section := range credentials.GetSectionNames() {
			credential := credentials.GetSection(section)

			result = append(result, *NewCredentialParamsFromValue(credential))
		}
	} else {
		credential := credentials.GetSection("credential")
		result = append(result, *NewCredentialParamsFromValue(credential))
	}

	return &result
}

func (c *CredentialParams) UseCredentialStore() bool {
	return c.GetAsNullableString("store_key") != nil
}

func (c *CredentialParams) StoreKey() string {
	return *c.GetAsNullableString("store_key")
}

func (c *CredentialParams) SetStoreKey(value string) {
	c.Put("store_key", value)
}

func (c *CredentialParams) Username() string {
	username := *c.GetAsNullableString("username")
	if username == "" {
		username = *c.GetAsNullableString("user")
	}
	return username
}

func (c *CredentialParams) SetUsername(value string) {
	c.Put("username", value)
}

func (c *CredentialParams) Password() string {
	password := *c.GetAsNullableString("password")
	if password == "" {
		password = *c.GetAsNullableString("pass")
	}
	return password
}

func (c *CredentialParams) SetPassword(value string) {
	c.Put("password", value)
}

func (c *CredentialParams) AccessId() string {
	accessId := *c.GetAsNullableString("access_id")
	if accessId == "" {
		accessId = *c.GetAsNullableString("client_id")
	}
	return accessId
}

func (c *CredentialParams) SetAccessId(value string) {
	c.Put("access_id", value)
}

func (c *CredentialParams) AccessKey() string {
	accessKey := *c.GetAsNullableString("access_key")
	if accessKey == "" {
		accessKey = *c.GetAsNullableString("client_key")
	}
	return accessKey
}

func (c *CredentialParams) SetAccessKey(value string) {
	c.Put("access_key", value)
}

func NewCredentialParamsFromString(line string) *CredentialParams {
	return &CredentialParams{
		StringValueMap: *data.NewStringValueMapFromString(line),
	}
}

func NewManyCredentialParamsFromConfig(config *config.ConfigParams) []*CredentialParams {
	result := []*CredentialParams{}

	credentials := config.GetSection("credentials")

	if credentials.StringValueMap.Len() > 0 {
		for _, section := range credentials.GetSectionNames() {
			credential := credentials.GetSection(section)
			result = append(result, NewCredentialParams(credential.StringValueMap.Value()))
		}
	} else {
		credential := config.GetSection("credential")
		if credential.StringValueMap.Len() > 0 {
			result = append(result, NewCredentialParams(credential.StringValueMap.Value()))
		}
	}

	return result
}

func NewCredentialParamsFromConfig(config config.ConfigParams) *CredentialParams {
	credentials := NewManyCredentialParamsFromConfig(&config)
	if len(credentials) > 0 {
		return credentials[0]
	} else {
		return nil
	}
}
