package cast

func Int16(v interface{}) (int16, error) {

	switch value := v.(type) {
	case uint8:
		return int16(value), nil
	case int8:
		return int16(value), nil
	case int16:
		return int16(value), nil
	case int16er:
		return value.Int16()
	default:
		return 0, internalCannotCastComplainer{expectedType:"int16", actualType:typeof(value)}
	}
}

// MustInt16 is like Int16, expect panic()s on an error.
func MustInt16(v interface{}) int16 {

	x, err := Int16(v)
	if nil != err {
		panic(err)
	}

	return x
}

type int16er interface {
	Int16() (int16, error)
}
