package cast

import (
	"fmt"
)

// String converts `v` into a string, if it is of type:
//
//	[]byte
//
//	rune
//
//	[]rune
//
// 	string
//
// ... or of a type that fits:
//
//	interface {
//		String() (string, error)
//	}
//
// ... or of a type that fits:
//
//	interface {
//		String() string
//	}
//
// (This final interface matches fmt.Stringer.)
func String(v interface{}) (string, error) {

	switch value := v.(type) {
	case stringer:
		return value.String()
	case fmt.Stringer:
		return value.String(), nil
	case []byte:
		return string(value), nil
	case rune:
		return string(value), nil
	case []rune:
		return string(value), nil
	case string:
		return string(value), nil
	default:
		return "", internalCannotCastComplainer{expectedType:"string", actualType:typeof(value)}
	}
}

// MustString is like String, expect panic()s on an error.
func MustString(v interface{}) string {

	x, err := String(v)
	if nil != err {
		panic(err)
	}

	return x
}

type stringer interface {
	String() (string, error)
}
