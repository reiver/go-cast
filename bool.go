package cast

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
