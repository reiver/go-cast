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
//		String() (string, bool)
//	}
//
// ... or of a type that fits:
//
//	interface {
//		String() string
//	}
//
// (This final interface matches fmt.Stringer.)
func String(v any) (string, error) {

	switch value := v.(type) {
	case errStringer:
		return value.String()
	case boolStringer:
		x, casted := value.String()
		if !casted {
			return "", errNotCasted
		}
		return x, nil
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
func MustString(v any) string {

	x, err := String(v)
	if nil != err {
		panic(err)
	}

	return x
}

type errStringer interface {
	String() (string, error)
}

type boolStringer interface {
	String() (string, bool)
}
