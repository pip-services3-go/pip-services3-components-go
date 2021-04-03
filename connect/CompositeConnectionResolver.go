package connect

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/auth"
)

type ICompositeConnectionResolverOverrides interface {
	ValidateConnection(correlationId string, connection *ConnectionParams) error

	ValidateCredential(correlationId string, credential *auth.CredentialParams) error

	ComposeOptions(connections []*ConnectionParams, credential *auth.CredentialParams, parameters *config.ConfigParams) *config.ConfigParams

	MergeConnection(options *config.ConfigParams, connection *ConnectionParams) *config.ConfigParams

	MergeCredential(options *config.ConfigParams, credential *auth.CredentialParams) *config.ConfigParams

	MergeOptional(options *config.ConfigParams, parameters *config.ConfigParams) *config.ConfigParams

	FinalizeOptions(options *config.ConfigParams) *config.ConfigParams
}

/**
 * Helper class that resolves connection and credential parameters,
 * validates them and generates connection options.
 *
 *  ### Configuration parameters ###
 *
 * - connection(s):
 *   - discovery_key:               (optional) a key to retrieve the connection from [IDiscovery]]
 *   - protocol:                    communication protocol
 *   - host:                        host name or IP address
 *   - port:                        port number
 *   - uri:                         resource URI or connection string with all parameters in it
 * - credential(s):
 *   - store_key:                   (optional) a key to retrieve the credentials from [ICredentialStore]]
 *   - username:                    user name
 *   - password:                    user password
 *
 * ### References ###
 *
 * - *:discovery:*:*:1.0          (optional) [IDiscovery]] services to resolve connections
 * - *:credential-store:*:*:1.0   (optional) Credential stores to resolve credentials
 */
type CompositeConnectionResolver struct {
	Overrides ICompositeConnectionResolverOverrides

	// The connection options
	Options *config.ConfigParams

	// The connections resolver.
	ConnectionResolver *ConnectionResolver

	// The credentials resolver.
	CredentialResolver *auth.CredentialResolver

	// The cluster support (multiple connections)
	ClusterSupported bool

	// The default protocol
	DefaultProtocol string

	// The default port number
	DefaultPort int

	// The list of supported protocols
	SupportedProtocols []string
}

// InheritCompositeConnectionResolver creates new CompositeConnectionResolver
// Parameters:
//   - overrides a child reference with overrides for virtual methods
// return *CompositeConnectionResolver
func InheritCompositeConnectionResolver(overrides ICompositeConnectionResolverOverrides) *CompositeConnectionResolver {
	return &CompositeConnectionResolver{
		Overrides:          overrides,
		ConnectionResolver: NewEmptyConnectionResolver(),
		CredentialResolver: auth.NewEmptyCredentialResolver(),
		ClusterSupported:   true,
		DefaultPort:        0,
	}
}

// Configures component by passing configuration parameters.
// - config    configuration parameters to be set.
func (c *CompositeConnectionResolver) Configure(config *config.ConfigParams) {
	c.ConnectionResolver.Configure(config)
	c.CredentialResolver.Configure(config)
	c.Options = config.GetSection("options")
}

// Sets references to dependent components.
// - references 	references to locate the component dependencies.
func (c *CompositeConnectionResolver) SetReferences(references refer.IReferences) {
	c.ConnectionResolver.SetReferences(references)
	c.CredentialResolver.SetReferences(references)
}

//  Resolves connection options from connection and credential parameters.
//  - correlationId     (optional) transaction id to trace execution through call chain.
//  - return 			 resolved options or error.
func (c *CompositeConnectionResolver) Resolve(correlationId string) (options *config.ConfigParams, err error) {
	var connections []*ConnectionParams
	var credential *auth.CredentialParams

	connections, err = c.ConnectionResolver.ResolveAll(correlationId)

	// Validate if cluster (multiple connections) is supported
	if err == nil && len(connections) > 0 && !c.ClusterSupported {
		err = cerr.NewConfigError(
			correlationId,
			"MULTIPLE_CONNECTIONS_NOT_SUPPORTED",
			"Multiple (cluster) connections are not supported",
		)
	}

	// Validate connections
	if err == nil {
		for _, connection := range connections {
			err = c.ValidateConnection(correlationId, connection)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		return nil, err
	}

	credential, err = c.CredentialResolver.Lookup(correlationId)
	if credential == nil {
		credential = auth.NewEmptyCredentialParams()
	}
	// Validate credential
	if err == nil {
		err = c.ValidateCredential(correlationId, credential)
	}

	if err != nil {
		return nil, err
	}

	return c.ComposeOptions(connections, credential, c.Options), nil
}

// Composes Composite connection options from connection and credential parameters.
// - correlationId     (optional) transaction id to trace execution through call chain.
// - connections       connection parameters
// - credential        credential parameters
// - parameters        optional parameters
// - return 		   resolved options or error.
func (c *CompositeConnectionResolver) Compose(correlationId string, connections []*ConnectionParams, credential *auth.CredentialParams,
	parameters *config.ConfigParams) (options *config.ConfigParams, err error) {

	// Validate connection parameters
	for _, connection := range connections {
		err = c.Overrides.ValidateConnection(correlationId, connection)
		if err != nil {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	// Validate credential parameters
	err = c.Overrides.ValidateCredential(correlationId, credential)

	if err != nil {
		return nil, err
	}

	return c.Overrides.ComposeOptions(connections, credential, parameters), nil
}

// Validates connection parameters.
// This method can be overriden in child classes.
// - correlationId    (optional) transaction id to trace execution through call chain.
// - connection       connection parameters to be validated
// - returns          error or nil if validation was successful
func (c *CompositeConnectionResolver) ValidateConnection(correlationId string, connection *ConnectionParams) error {
	if connection == nil {
		return cerr.NewConfigError(correlationId, "NO_CONNECTION", "Connection parameters are not set is not set")
	}

	// URI usually contains all information
	uri := connection.Uri()
	if uri != "" {
		return nil
	}

	protocol := connection.ProtocolWithDefault(c.DefaultProtocol)
	if protocol == "" {
		return cerr.NewConfigError(correlationId, "NO_PROTOCOL", "Connection protocol is not set")
	}
	if c.SupportedProtocols != nil && indexOf(c.SupportedProtocols, protocol) < 0 {
		return cerr.NewConfigError(correlationId, "UNSUPPORTED_PROTOCOL", "The protocol "+protocol+" is not supported")
	}

	var host = connection.Host()
	if host == "" {
		return cerr.NewConfigError(correlationId, "NO_HOST", "Connection host is not set")
	}

	var port = connection.PortWithDefault(c.DefaultPort)
	if port == 0 {
		return cerr.NewConfigError(correlationId, "NO_PORT", "Connection port is not set")
	}

	return nil
}

// Validates credential parameters.
// This method can be overriden in child classes.
// - correlationId     (optional) transaction id to trace execution through call chain.
// - credential  credential parameters to be validated
// - returns error or nil if validation was successful
func (c *CompositeConnectionResolver) ValidateCredential(correlationId string, credential *auth.CredentialParams) error {
	return nil
}

// Composes connection and credential parameters into connection options.
// This method can be overriden in child classes.
// - connections a list of connection parameters
// - credential credential parameters
// - parameters optional parameters
// - returns a composed connection options.
func (c *CompositeConnectionResolver) ComposeOptions(connections []*ConnectionParams, credential *auth.CredentialParams, parameters *config.ConfigParams) *config.ConfigParams {
	// Connection options
	options := config.NewEmptyConfigParams()

	// Merge connection parameters
	for _, connection := range connections {
		options = c.Overrides.MergeConnection(options, connection)
	}

	// Merge credential parameters
	options = c.Overrides.MergeCredential(options, credential)

	// Merge optional parameters
	options = c.Overrides.MergeOptional(options, parameters)

	// Perform final processing
	options = c.Overrides.FinalizeOptions(options)

	return options
}

// Merges connection options with connection parameters
// This method can be overriden in child classes.
// -  options connection options
// -  connection connection parameters to be merged
// - returns merged connection options.
func (c *CompositeConnectionResolver) MergeConnection(options *config.ConfigParams, connection *ConnectionParams) *config.ConfigParams {
	var mergedOptions = options.SetDefaults(&connection.ConfigParams)
	return mergedOptions
}

// Merges connection options with credential parameters
// This method can be overriden in child classes.
// - options connection options
// - credential credential parameters to be merged
// - returns merged connection options.
func (c *CompositeConnectionResolver) MergeCredential(options *config.ConfigParams, credential *auth.CredentialParams) *config.ConfigParams {
	var mergedOptions = options.Override(&credential.ConfigParams)
	return mergedOptions
}

// Merges connection options with optional parameters
// This method can be overriden in child classes.
// - options connection options
// - parameters optional parameters to be merged
// - returns merged connection options.
func (c *CompositeConnectionResolver) MergeOptional(options *config.ConfigParams, parameters *config.ConfigParams) *config.ConfigParams {
	var mergedOptions = options.SetDefaults(parameters)
	return mergedOptions
}

// Finalize merged options
// This method can be overriden in child classes.
// - options connection options
// - returns finalized connection options
func (c *CompositeConnectionResolver) FinalizeOptions(options *config.ConfigParams) *config.ConfigParams {
	return options
}

func indexOf(a []string, e string) int {
	for i := range a {
		if e == a[i] {
			return i
		}
	}
	return -1
}
