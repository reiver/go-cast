package cast

import (
	"fmt"
)

// The CannotSetComplainer interface identifies an error due to a situation
// where setting a value of some particular type cannot happen.
//
// This can likely be thought of some type of bug in the code,
// due to the destination-type or the source-type being the wrong type.
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
type CannotSetComplainer interface {
	error
	CannotSetComplainer()
}

type internalCannotSetComplainer struct {
	srcType string
	dstType string
}

func (receiver internalCannotSetComplainer) Error() string {
	return fmt.Sprintf("cast: cannot set something of type %q into something of type %q", receiver.srcType, receiver.dstType)
}

func (internalCannotSetComplainer) CannotSetComplainer() {
	// Nothing here.
}

var _ CannotSetComplainer = internalCannotSetComplainer{}
