package cast

// Bool will return a bool when `v` is of type bool, or has a method:
//
//	type interface {
//		Bool() (bool, error)
//	}
//
// ... that returns successfully.
//
// Else it will return an error.
func Bool(v any) (bool, error) {

	switch value := v.(type) {
	case bool:
		return bool(value), nil
	case booler:
		return value.Bool()
	default:
		return false, internalCannotCastComplainer{expectedType:"bool", actualType:typeof(value)}
	}
}

// BoolElse is similar to [Bool] except that if a cast cannot be done, it returns the `alternative`.
func BoolElse(v any, alternative bool) bool {
	result, err := Bool(v)
	if nil != err {
		return alternative
	}
	return result
}

// MustBool is like Bool, expect panic()s on an error.
func MustBool(v any) bool {

	x, err := Bool(v)
	if nil != err {
		panic(err)
	}

	return x
}

type booler interface {
	Bool() (bool, error)
}
