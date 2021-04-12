package components

import (
	"github.com/pip-services3-go/pip-services3-commons-go/config"
	"github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/pip-services3-go/pip-services3-components-go/count"
	"github.com/pip-services3-go/pip-services3-components-go/log"
)

type Component struct {
	DependencyResolver *refer.DependencyResolver
	Logger             *log.CompositeLogger
	Counters           *count.CompositeCounters
}

func InheritComponent() *Component {
	return &Component{
		DependencyResolver: refer.NewDependencyResolver(),
		Logger:             log.NewCompositeLogger(),
		Counters:           count.NewCompositeCounters(),
	}
}

func (c *Component) Configure(config *config.ConfigParams) {
	c.DependencyResolver.Configure(config)
	c.Logger.Configure(config)
}

func (c *Component) SetReferences(references refer.IReferences) {
	c.DependencyResolver.SetReferences(references)
	c.Logger.SetReferences(references)
	c.Counters.SetReferences(references)
}
