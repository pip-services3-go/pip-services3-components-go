package connect

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
)

type ConnectionResolver struct {
	_connections []*ConnectionParams
	_references  refer.IReferences
}

func NewConnectionResolver(conf *config.ConfigParams, references refer.IReferences) *ConnectionResolver {
	result := &ConnectionResolver{}
	if conf != nil {
		result.Configure(conf)
	}
	if references != nil {
		result.SetReferences(references)
	}
	return result
}

func (conR *ConnectionResolver) Configure(conf *config.ConfigParams) {
	conR._connections = NewManyConnectionParamsFromConfig(conf)
}

func (conR *ConnectionResolver) SetReferences(references refer.IReferences) {
	conR._references = references
}

func (conR ConnectionResolver) GetAll() []*ConnectionParams {
	conns := make([]*ConnectionParams, len(conR._connections))
	for i := range conR._connections {
		conns[i] = NewConnectionParams(conR._connections[i].Value())
	}
	return conns
}

func (conR *ConnectionResolver) Add(conn *ConnectionParams) {
	conR._connections = append(conR._connections, NewConnectionParams(conn.Value()))
}

func (conR *ConnectionResolver) resolveInDiscovery(correlationId string, connection *ConnectionParams) (resolvedConnectionParams *ConnectionParams, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			resolvedConnectionParams = nil
		}
	}()
	if !connection.UseDiscovery() {
		return
	}

	if conR._references == nil {
		return
	}

	discoveryDescriptor := refer.NewDescriptor("*", "discovery", "*", "*", "*")
	discoveries := conR._references.GetOptional(discoveryDescriptor)
	if len(discoveries) == 0 {
		err = refer.NewReferenceError(correlationId, discoveryDescriptor)
		return
	}

	key := connection.GetDiscoveryKey()
	for i := range discoveries {
		var res *ConnectionParams = nil
		res, err = discoveries[i].(IDiscovery).ResolveOne(correlationId, key)
		if err == nil || res != nil {
			resolvedConnectionParams = NewConnectionParams(res.Value())
			break
		}
	}
	return
}

func (conR *ConnectionResolver) Resolve(correlationId string) (resolvedConnectionParams *ConnectionParams, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			resolvedConnectionParams = nil
		}
	}()

	if len(conR._connections) == 0 {
		return
	}

	connections := make([]*ConnectionParams, 0)

	for i := range conR._connections {
		if !conR._connections[i].UseDiscovery() {
			return conR._connections[i], nil
		} else {
			connections = append(connections, NewConnectionParams(conR._connections[i].Value()))
		}
	}
	if len(connections) == 0 {
		return
	}
	for i := range connections {
		var res *ConnectionParams = nil
		res, err = conR.resolveInDiscovery(correlationId, connections[i])
		if err == nil || res != nil {
			resolvedConnectionParams = NewConnectionParams(config.NewConfigParamsFromMaps(connections[i].Value(), res.Value()).Value())
			break
		}
	}
	return
}

func (conR *ConnectionResolver) resolveAllInDiscovery(correlationId string, connection *ConnectionParams) (resolvedConnectionsParams []*ConnectionParams, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			resolvedConnectionsParams = nil
		}
	}()

	resolvedConnectionsParams = make([]*ConnectionParams, 0)

	if !connection.UseDiscovery() {
		return
	}

	if conR._references == nil {
		return
	}

	discoveryDescriptor := refer.NewDescriptor("*", "discovery", "*", "*", "*")
	discoveries := conR._references.GetOptional(discoveryDescriptor)
	if len(discoveries) == 0 {
		err = refer.NewReferenceError(correlationId, discoveryDescriptor)
		return
	}

	key := connection.GetDiscoveryKey()

	for i := range discoveries {
		var res []*ConnectionParams = nil
		res, err = discoveries[i].(IDiscovery).ResolveAll(correlationId, key)
		if err == nil || res != nil {
			resolvedConnectionsParams = append(resolvedConnectionsParams, res...)
		}
	}
	return
}

func (conR *ConnectionResolver) ResolveAll(correlationId string) (resolvedConnectionsParams []*ConnectionParams, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			resolvedConnectionsParams = make([]*ConnectionParams, 0)
		}
	}()

	toResolve := make([]*ConnectionParams, 0)
	resolvedConnectionsParams = make([]*ConnectionParams, 0)

	for i := range conR._connections {
		if conR._connections[i].UseDiscovery() {
			toResolve = append(toResolve, NewConnectionParams(conR._connections[i].Value()))
		} else {
			resolvedConnectionsParams = append(resolvedConnectionsParams, NewConnectionParams(conR._connections[i].Value()))
		}
	}

	if len(toResolve) == 0 {
		return
	}

	for i := range toResolve {
		var res []*ConnectionParams = nil
		res, err = conR.resolveAllInDiscovery(correlationId, toResolve[i])
		if err != nil {
			break
		}
		for j := range res {
			resolvedConnectionsParams = append(resolvedConnectionsParams, NewConnectionParams(res[j].Value()))
		}
	}
	return
}

func (conR *ConnectionResolver) registerInDiscovery(correlationId string, connection *ConnectionParams) (ok bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			ok = false
		}
	}()
	if !connection.UseDiscovery() {
		return
	}

	if conR._references == nil {
		return
	}

	discoveryDescriptor := refer.NewDescriptor("*", "discovery", "*", "*", "*")
	discoveries := conR._references.GetOptional(discoveryDescriptor)
	if discoveries == nil {
		return
	}

	key := connection.GetDiscoveryKey()

	for i := range discoveries {
		_, err = discoveries[i].(IDiscovery).Register(correlationId, key, connection)
		if err != nil {
			return
		}
	}
	ok = true
	return
}

func (conR *ConnectionResolver) Register(correlationId string, key string, connection *ConnectionParams) (T interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	var ok bool
	if ok, err = conR.registerInDiscovery(correlationId, connection); ok {
		conR._connections = append(conR._connections, NewConnectionParams(connection.Value()))
	}
	return
}
