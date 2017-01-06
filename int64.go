package cast

func Int64(v interface{}) (int64, error) {

	switch value := v.(type) {
	case uint8:
		return int64(value), nil
	case uint16:
		return int64(value), nil
	case uint32:
		return int64(value), nil
	case int:
		return int64(value), nil
	case int8:
		return int64(value), nil
	case int16:
		return int64(value), nil
	case int32:
		return int64(value), nil
	case int64:
		return int64(value), nil
	case int64er:
		return value.Int64()
	default:
		return 0, internalCannotCastComplainer{expectedType:"int64", actualType:typeof(value)}
	}
}

func MustInt64(v interface{}) (int64, error) {

	x, err := Int64(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type int64er interface {
	Int64() (int64, error)
}
