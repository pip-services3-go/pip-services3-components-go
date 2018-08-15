package build

type IFactory interface {
	CanCreate(locator interface{}) interface{}
	Create(locator interface{}) (interface{}, error)
}
