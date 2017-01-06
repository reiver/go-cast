package cast

func Uint(v interface{}) (uint, error) {

	switch value := v.(type) {
	case uint8:
		return uint(value), nil
	case uint16:
		return uint(value), nil
	case uint32:
		return uint(value), nil
	case uinter:
		return value.Uint()
	default:
		return 0, internalCannotCastComplainer{expectedType:"uint", actualType:typeof(value)}
	}
}

func MustUint(v interface{}) (uint, error) {

	x, err := Uint(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type uinter interface {
	Uint() (uint, error)
}
