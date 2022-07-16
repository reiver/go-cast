package cast

import (
	"github.com/reiver/go-fck"
)

const (
	errNilPointer = fck.Error("nil pointer")
	errNotCasted  = fck.Error("not casted")
	errNotOK      = fck.Error("not ok")
)
