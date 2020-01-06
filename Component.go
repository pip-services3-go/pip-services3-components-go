package components

import (
	"github.com/pip-services3-go/pip-services3-commons-go/v3/config"
	"github.com/pip-services3-go/pip-services3-commons-go/v3/refer"
	"github.com/pip-services3-go/pip-services3-components-go/v3/count"
	"github.com/pip-services3-go/pip-services3-components-go/v3/log"
)

type Component struct {
	dependencyResolver *refer.DependencyResolver
	logger             *log.CompositeLogger
	counters           *count.CompositeCounters
}

func InheritComponent() *Component {
	return &Component{
		dependencyResolver: refer.NewDependencyResolver(),
		logger:             log.NewCompositeLogger(),
		counters:           count.NewCompositeCounters(),
	}
}

func (c *Component) Configure(config *config.ConfigMap) {
	c.dependencyResolver.Configure(config)
	c.logger.Configure(config)
}

func (c *Component) SetReferences(references refer.IReferences) {
	c.dependencyResolver.SetReferences(references)
	c.logger.SetReferences(references)
	c.counters.SetReferences(references)
}
