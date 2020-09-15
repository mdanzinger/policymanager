package policy

import (
	"errors"
	"fmt"
)

var (
	ErrUnexpectedActionInput    = errors.New("unexpected action value type")
	ErrUnexpectedConditionInput = errors.New("unexpected condition value type")

	// Todo: make this a custom error type
	ErrUnexpectedSpecFieldType = errors.New("unexpected spec field type")
)

type ErrEnforcerNotFound struct {
	policyType string
}

func (e ErrEnforcerNotFound) Error() string {
	return fmt.Sprintf("unable to find policyType for policy type %s", e.policyType)
}

func NewErrEnforcerNotFound(policyType string) error {
	return ErrEnforcerNotFound{policyType: policyType}
}
