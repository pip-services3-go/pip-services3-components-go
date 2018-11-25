package auth

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/errors"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
)

type CredentialResolver struct {
	_credentials []CredentialParams
	_references  refer.IReferences
}

func NewCredentialResolver(config config.ConfigParams, references refer.References) *CredentialResolver {
	resolver := CredentialResolver{}
	resolver.Configure(config)
	resolver.SetReferences(references)

	return &resolver
}

func (cr *CredentialResolver) Configure(config config.ConfigParams) {
	cr._credentials = append(cr._credentials, *NewCredentialParamsManyFromConfig(config)...)
}

func (cr *CredentialResolver) SetReferences(references refer.References) {
	cr._references = &references
}

func (cr *CredentialResolver) GetAll() []CredentialParams {
	return cr._credentials
}

func (cr *CredentialResolver) Add(credential CredentialParams) {
	cr._credentials = append(cr._credentials, credential)
}

func (cr *CredentialResolver) LookupInStores(correlationId string, credential CredentialParams) (*errors.ApplicationError, *CredentialParams) {
	if !credential.UseCredentialStore() {
		return nil, nil
	}

	if cr._references == nil {
		return nil, nil
	}

	storeDescriptor := refer.NewDescriptor("*", "credential-store", "*", "*", "*")

	components := cr._references.GetOptional(storeDescriptor)

	if len(components) == 0 {
		return refer.NewReferenceError(correlationId, storeDescriptor), nil
	}

	key := credential.StoreKey()

	for _, component := range components {
		store := component.(ICredentialStore)
		result, error := store.Lookup(correlationId, key)
		if error == nil && result != nil {
			return nil, result
		}
	}

	return nil, nil
}
