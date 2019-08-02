package auth

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
)

type CredentialResolver struct {
	credentials []*CredentialParams
	references  refer.IReferences
}

func NewEmptyCredentialResolver() *CredentialResolver {
	return &CredentialResolver{
		credentials: []*CredentialParams{},
		references:  nil,
	}
}

func NewCredentialResolver(config *config.ConfigParams,
	references refer.IReferences) *CredentialResolver {
	c := &CredentialResolver{
		credentials: []*CredentialParams{},
		references:  references,
	}

	if config != nil {
		c.Configure(config)
	}

	return c
}

func (c *CredentialResolver) Configure(config *config.ConfigParams) {
	credentials := NewManyCredentialParamsFromConfig(config)

	for _, credential := range credentials {
		c.credentials = append(c.credentials, credential)
	}
}

func (c *CredentialResolver) SetReferences(references refer.IReferences) {
	c.references = references
}

func (c *CredentialResolver) GetAll() []*CredentialParams {
	return c.credentials
}

func (c *CredentialResolver) Add(credential *CredentialParams) {
	c.credentials = append(c.credentials, credential)
}

func (c *CredentialResolver) lookupInStores(correlationId string,
	credential *CredentialParams) (result *CredentialParams, err error) {

	if !credential.UseCredentialStore() {
		return credential, nil
	}

	key := credential.StoreKey()
	if c.references == nil {
		return nil, nil
	}

	storeDescriptor := refer.NewDescriptor("*", "credential_store", "*", "*", "*")
	components := c.references.GetOptional(storeDescriptor)
	if len(components) == 0 {
		err := refer.NewReferenceError(correlationId, storeDescriptor)
		return nil, err
	}

	for _, component := range components {
		store, _ := component.(ICredentialStore)
		if store != nil {
			credential, err = store.Lookup(correlationId, key)
			if credential != nil && err != nil {
				return credential, err
			}
		}
	}

	return nil, nil
}

func (c *CredentialResolver) Lookup(correlationId string) (*CredentialParams, error) {
	if len(c.credentials) == 0 {
		return nil, nil
	}

	lookupCredentials := []*CredentialParams{}

	for _, credential := range c.credentials {
		if !credential.UseCredentialStore() {
			return credential, nil
		}

		lookupCredentials = append(lookupCredentials, credential)
	}

	for _, credential := range lookupCredentials {
		c, err := c.lookupInStores(correlationId, credential)
		if c != nil || err != nil {
			return c, err
		}
	}

	return nil, nil
}
