package cast

// Bool will return a bool when `v` is of type bool, or has a method:
//
//	type interface {
//		Bool() (bool, error)
//	}
//
// .. that returns successfully.
//
// Else it will return an error.
func Bool(v interface{}) (bool, error) {

	switch value := v.(type) {
	case bool:
		return bool(value), nil
	case booler:
		return value.Bool()
	default:
		return false, internalCannotCastComplainer{expectedType:"bool", actualType:typeof(value)}
	}
}

// MustBool is like Bool, expect panic()s on an error.
func MustBool(v interface{}) bool {

	x, err := Bool(v)
	if nil != err {
		panic(err)
	}

	return x
}

type booler interface {
	Bool() (bool, error)
}
