package cast

func Int16(v interface{}) (int16, error) {

	switch value := v.(type) {
	case uint8:
		return int16(value), nil
	case int8:
		return int16(value), nil
	case int16:
		return int16(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"int16", actualType:typeof(value)}
	}
}

func MustInt16(v interface{}) (int16, error) {

	x, err := Int16(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}
