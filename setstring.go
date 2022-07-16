package cast

import (
	"database/sql"
	"fmt"
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
		case stringSetter:
			casted.SetString(value)
		case errableStringSetter:
			e := casted.SetString(value)
			if nil != e {
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:"interface{SetString(string)error}",
					err:err,
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
		case flagStringSetter:
			e := casted.Set(value)
			if nil != e {
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:"interface{Set(string)error}",
					err:err,
				}
			}
		case sql.Scanner:
			e := casted.Scan(value)
			if nil != e {
				err = internalCouldNotSetComplainer{
					srcType:srcType,
					dstType:"interface{Scan(any)error}",
					err:err,
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
