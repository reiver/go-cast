package cast

import (
	"fmt"
)

type CannotCastComplainer interface {
	error
	CannotCastComplainer()
}

type internalCannotCastComplainer struct{
	expectedType string
	actualType string
}

func (receiver internalCannotCastComplainer) Error() string {
	return fmt.Sprintf("Cannot Cast: %q; expected something compatible with: %q", receiver.actualType, receiver.expectedType)
}

func (internalCannotCastComplainer) CannotCastComplainer() {
	// Nothing here.
}
