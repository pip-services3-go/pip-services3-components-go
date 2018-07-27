package build

import (
	"fmt"

	"github.com/pip-services/pip-services-commons-go/errors"
)

func NewCreateError(correlationId string, locator interface{}) *ApplicationError {
	e := errors.NewInternalError(correlationId, "CANNOT_CREATE", message)
	return e
}

func NewCreateErrorByLocator(correlationId string, locator interface{}) *ApplicationError {
	message := fmt.Sprintf("Requested component %v cannot be created", locator)
	e := errors.NewInternalError(correlationId, "CANNOT_CREATE", message)
	e.WithDetails("locator", locator)
	return e
}
