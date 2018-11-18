package build

import (
	"fmt"

	"github.com/pip-services3-go/pip-services3-commons-go/errors"
)

func NewCreateError(correlationId string, message string) *errors.ApplicationError {
	e := errors.NewInternalError(correlationId, "CANNOT_CREATE", message)
	return e
}

func NewCreateErrorByLocator(correlationId string, locator interface{}) *errors.ApplicationError {
	message := fmt.Sprintf("Requested component %v cannot be created", locator)
	e := errors.NewInternalError(correlationId, "CANNOT_CREATE", message)
	e.WithDetails("locator", locator)
	return e
}
