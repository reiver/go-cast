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

func MustBool(v interface{}) (bool, error) {

	x, err := Bool(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type booler interface {
	Bool() (bool, error)
}
