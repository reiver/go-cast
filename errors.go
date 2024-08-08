package cast

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilPointer = erorr.Error("cast: nil pointer")
	errNotCasted  = erorr.Error("cast: not casted")
	errNotOK      = erorr.Error("cast: not ok")
)
