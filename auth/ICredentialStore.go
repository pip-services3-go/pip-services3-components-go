package auth

type ICredentialStore interface {
	Store(correlationId string, key string, credential *CredentialParams) error
	Lookup(correlationId string, key string) (*CredentialParams, error)
}
