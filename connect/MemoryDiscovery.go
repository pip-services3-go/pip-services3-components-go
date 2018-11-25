package connect

import "github.com/pip-services3-go/pip-services3-commons-go/config"

type DiscoveryItem struct {
	Key        *string
	Connection *ConnectionParams
}

type MemoryDiscovery struct {
	_items []*DiscoveryItem
}

func NewMemoryDiscovery(conf *config.ConfigParams) (memDisc *MemoryDiscovery) {
	memDisc = &MemoryDiscovery{}
	if conf != nil {
		memDisc.Configure(conf)
	}
	return
}

func (md *MemoryDiscovery) Configure(conf *config.ConfigParams) {
	md.ReadConnections(conf)
}

func (md *MemoryDiscovery) ReadConnections(conf *config.ConfigParams) {
	keys := conf.GetSectionNames()
	md._items = make([]*DiscoveryItem, len(keys))
	for i, key := range keys {
		md._items[i] = &DiscoveryItem{
			Key:        &key,
			Connection: NewConnectionParamsFromString(*conf.GetAsNullableString(key)),
		}
	}
}

func (md *MemoryDiscovery) Register(correlationId string, key string, connection *ConnectionParams) (res interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	md._items = append(md._items, &DiscoveryItem{
		Key:        &key,
		Connection: NewConnectionParams(connection.Value()),
	})
	return
}

func (md *MemoryDiscovery) ResolveOne(correlationId string, key string) (conParams *ConnectionParams, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			conParams = nil
		}
	}()

	for _, item := range md._items {
		if *item.Key == key && item.Connection != nil {
			conParams = NewConnectionParams(item.Connection.Value())
			break
		}
	}
	return
}

func (md *MemoryDiscovery) ResolveAll(correlationId string, key string) (consParams []*ConnectionParams, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			consParams = nil
		}
	}()
	consParams = make([]*ConnectionParams, 0)
	for _, item := range md._items {
		if *item.Key == key && item.Connection != nil {
			consParams = append(consParams, NewConnectionParams(item.Connection.Value()))
		}
	}
	return
}
