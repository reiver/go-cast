package cast

import (
	"fmt"
)

// The CannotCastComplainer interface identifies an error due to a situation
// where converting from one type to another could not be done.
//
// Example:
//
//	s, err := cast.String(v)
//	if nil != err {
//		switch {
//		case cast.CannotCastComplainer:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
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
