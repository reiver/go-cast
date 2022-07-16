package cast

import (
	"database/sql"
	"fmt"
	"io"
)

// SetString uses cast.String() to try to convert `src` into a string, and then (if successful) tries to assign that string to `dst` if `dst` is of type:
//
//	*string
//
//	*[]byte
//
//	*[]rune
//
// ... or of a type that fits:
//
//	interface {
//		SetString(string)
//	}
// ... or of a type that fits:
//
//	interface {
//		SetString(string) error
//	}
//
// ... or of a type that fits:
//
//	interface {
//		SetString(string) bool
//	}
//
// ... or of a type that fits:
//
//	interface {
//		Set(string) error
//	}
//
// (This interface is part of flag.Value.)
//
// ... or of a type that fits:
//
//	interface {
//		Scan(any) error
//	}
//
// (This interface matches sql.Scanner.)
func SetString(dst any, src any) error {

	var value string
	{
		var err error

		value, err = String(src)
		if nil != err {
			return err
		}
	}

	var err error
	{
		const srcType = "string"

		switch casted := dst.(type) {
		case errableStringSetter:
			e := casted.SetString(value)
			if nil != e {
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:"interface{SetString(string)error}",
					err:e,
				}
			}
		case fallibleStringSetter:
			ok := casted.SetString(value)
			if !ok {
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:"interface{SetString(string)bool}",
					err:errNotOK,
				}
			}
		case stringSetter:
			casted.SetString(value)
		case io.Writer:
			n, e := io.WriteString(casted, value)
			switch {
			case nil != err:
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:"interface{Write([]byte)(int,error)}",
					err:e,
				}
			case n != len(value):
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:"interface{Write([]byte)(int,error)}",
					err:fmt.Errorf("expected to write %d bytes but actually wrote %d bytes", len(value), n),
				}
			}
		case flagStringSetter:
			e := casted.Set(value)
			if nil != e {
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:"interface{Set(string)error}",
					err:e,
				}
			}
		case sql.Scanner:
			e := casted.Scan(value)
			if nil != e {
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:"interface{Scan(any)error}",
					err:e,
				}
			}
		case *string:
			switch {
			case nil == casted:
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:fmt.Sprintf("%T",casted),
					err:errNilPointer,
				}
			default:
				*casted = value
			}
		case *[]byte:
			switch {
			case nil == casted:
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:fmt.Sprintf("%T",casted),
					err:errNilPointer,
				}
			default:
				*casted = []byte(value)
			}
		case *[]rune:
			switch {
			case nil == casted:
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:fmt.Sprintf("%T",casted),
					err:errNilPointer,
				}
			default:
				*casted = []rune(value)
			}
		default:
			err = internalCannotSetComplainer{
				srcType:srcType,
				dstType:fmt.Sprintf("%T",dst),
			}
		}
	}

	return err
}

type stringSetter interface {
	SetString(string)
}

type errableStringSetter interface {
	SetString(string)error
}

type fallibleStringSetter interface {
	SetString(string)bool
}

type flagStringSetter interface {
	Set(string) error
}
