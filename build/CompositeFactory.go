package build

type CompositeFactory struct {
	factories []IFactory
}

func NewCompositeFactory() *CompositeFactory {
	return &CompositeFactory{
		factories: []IFactory{},
	}
}

func NewCompositeFactoryFromFactories(factories ...IFactory) *CompositeFactory {
	return &CompositeFactory{
		factories: factories,
	}
}

func (c *CompositeFactory) Add(factory IFactory) {
	if factory == nil {
		panic("Factory cannot be nil")
	}

	c.factories = append(c.factories, factory)
}

func (c *CompositeFactory) Remove(factory IFactory) {
	for i, thisFactory := range c.factories {
		if thisFactory == factory {
			c.factories = append(c.factories[:i], c.factories[i+1:]...)
			break
		}
	}
}

func (c *CompositeFactory) CanCreate(locator interface{}) interface{} {
	if locator == nil {
		panic("Locator cannot be null")
	}

	// Iterate from the latest factories
	for i := len(c.factories) - 1; i >= 0; i-- {
		thisLocator := c.factories[i].CanCreate(locator)
		if thisLocator != nil {
			return thisLocator
		}
	}

	return nil
}

func (c *CompositeFactory) Create(locator interface{}) (interface{}, error) {
	if locator == nil {
		panic("Locator cannot be null")
	}

	// Iterate from the latest factories
	for i := len(c.factories) - 1; i >= 0; i-- {
		factory := c.factories[i]
		if factory.CanCreate(locator) != nil {
			return factory.Create(locator)
		}
	}

	return nil, NewCreateErrorByLocator("", locator)
}
