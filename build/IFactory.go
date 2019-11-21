package build

/*
Interface for component factories.

Factories use locators to identify components to be created.

The locators are similar to those used to locate components in references. They can be of any type like strings or integers. However Pip.Services toolkit most often uses Descriptor objects as component locators.
*/
type IFactory interface {
	// Checks if this factory is able to create component by given locator.
	// This method searches for all registered components and returns a locator for component it is able to create that matches the given locator. If the factory is not able to create a requested component is returns null.
	CanCreate(locator interface{}) interface{}
	// 	Creates a component identified by given locator.
	// Rturn a CreateError if the factory is not able to create the component.
	Create(locator interface{}) (interface{}, error)
}
