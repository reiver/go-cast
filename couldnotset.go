package cast

import (
	"fmt"
)

// The CouldNotSetComplainer interface identifies an error due to a situation
// where setting a value of some particular type failed with some type of error.
//
// This can likely be thought of some type of run-time error.
//
// Example:
//
//	err := cast.SetString(v)
//	if nil != err {
//		switch {
//		case cast.CannotSetComplainer:
//			//@TODO
//		case cast.CouldNotSetComplainer:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type CouldNotSetComplainer interface {
	error
	CouldNotSetComplainer()
	Unwrap() error
}

type internalCouldNotSetComplainer struct {
	srcType string
	dstType string
	err     error
}

func (receiver internalCouldNotSetComplainer) Error() string {
	return fmt.Sprintf("cast: could not set something of type %q into something of type %q because of error: %s", receiver.srcType, receiver.dstType, receiver.err)
}

func (internalCouldNotSetComplainer) CouldNotSetComplainer() {
	// Nothing here.
}

func (receiver internalCouldNotSetComplainer) Unwrap() error {
	return receiver.err
}

var _ CouldNotSetComplainer = internalCouldNotSetComplainer{}
