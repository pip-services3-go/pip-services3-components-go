package auth

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
)

type MemoryCredentialStore struct {
	items map[string]*CredentialParams
}

func NewEmptyMemoryCredentialStore() *MemoryCredentialStore {
	return &MemoryCredentialStore{
		items: map[string]*CredentialParams{},
	}
}

func NewMemoryCredentialStore(config *config.ConfigParams) *MemoryCredentialStore {
	c := &MemoryCredentialStore{
		items: map[string]*CredentialParams{},
	}

	if config != nil {
		c.Configure(config)
	}

	return c
}

func (c *MemoryCredentialStore) Configure(config *config.ConfigParams) {
	c.ReadCredentials(config)
}

func (c *MemoryCredentialStore) ReadCredentials(config *config.ConfigParams) {
	c.items = map[string]*CredentialParams{}

	keys := config.Keys()
	for _, key := range keys {
		value := config.GetAsString(key)
		credential := NewCredentialParamsFromString(value)
		c.items[key] = credential
	}
}

func (c *MemoryCredentialStore) Store(correlationId string, key string,
	credential *CredentialParams) error {

	if credential != nil {
		c.items[key] = credential
	} else {
		delete(c.items, key)
	}

	return nil
}

func (c *MemoryCredentialStore) Lookup(correlationId string,
	key string) (result *CredentialParams, err error) {
	credential, _ := c.items[key]
	return credential, nil
}
