package cast

func Int32(v interface{}) (int32, error) {

	switch value := v.(type) {
	case uint8:
		return int32(value), nil
	case uint16:
		return int32(value), nil
	case int8:
		return int32(value), nil
	case int16:
		return int32(value), nil
	case int32:
		return int32(value), nil
	case int32er:
		return value.Int32()
	default:
		return 0, internalCannotCastComplainer{expectedType:"int32", actualType:typeof(value)}
	}
}

func MustInt32(v interface{}) (int32, error) {

	x, err := Int32(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type int32er interface {
	Int32() (int32, error)
}
