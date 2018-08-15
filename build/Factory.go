package build

import (
	refl "reflect"

	"github.com/pip-services-go/pip-services-commons-go/convert"
	"github.com/pip-services-go/pip-services-commons-go/data"
)

type registration struct {
	locator interface{}
	factory func() interface{}
}

type Factory struct {
	registrations []*registration
}

func NewFactory() *Factory {
	return &Factory{
		registrations: []*registration{},
	}
}

func (c *Factory) Register(locator interface{}, factory func() interface{}) {
	if locator == nil {
		panic("Locator cannot be nil")
	}
	if factory == nil {
		panic("Factory cannot be nil")
	}

	c.registrations = append(c.registrations, &registration{
		locator: locator,
		factory: factory,
	})
}

func (c *Factory) RegisterType(locator interface{}, factory interface{}) {
	if locator == nil {
		panic("Locator cannot be nil")
	}
	if factory == nil {
		panic("Factory cannot be nil")
	}

	val := refl.ValueOf(factory)
	if val.Kind() != refl.Func {
		panic("Factory must be parameterless function")
	}

	c.Register(locator, func() interface{} {
		return val.Call([]refl.Value{})[0].Interface()
	})
}

func (c *Factory) CanCreate(locator interface{}) interface{} {
	for _, registration := range c.registrations {
		thisLocator := registration.locator

		equatable, ok := thisLocator.(data.IEquatable)
		if ok && equatable.Equals(locator) {
			return thisLocator
		}

		if thisLocator == locator {
			return thisLocator
		}
	}
	return nil
}

func (c *Factory) Create(locator interface{}) (interface{}, error) {
	var factory func() interface{}

	for _, registration := range c.registrations {
		thisLocator := registration.locator

		equatable, ok := thisLocator.(data.IEquatable)
		if ok && equatable.Equals(locator) {
			factory = registration.factory
			break
		}

		if thisLocator == locator {
			factory = registration.factory
			break
		}
	}

	if factory == nil {
		return nil, NewCreateErrorByLocator("", locator)
	}

	var err error

	obj := func() interface{} {
		defer func() {
			if r := recover(); r != nil {
				tempMessage := convert.StringConverter.ToString(r)
				tempError := NewCreateError("", tempMessage)

				cause, ok := r.(error)
				if ok {
					tempError.WithCause(cause)
				}

				err = tempError
			}
		}()

		return factory()
	}()

	return obj, err
}
