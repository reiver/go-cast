package cast

func Int8(v interface{}) (int8, error) {

	switch value := v.(type) {
	case int8:
		return int8(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"int8", actualType:typeof(value)}
	}
}

func MustInt8(v interface{}) (int8, error) {

	x, err := Int8(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}
