package cast

func Int(v interface{}) (int, error) {

	switch value := v.(type) {
	case uint8:
		return int(value), nil
	case uint16:
		return int(value), nil
	case int:
		return int(value), nil
	case int8:
		return int(value), nil
	case int16:
		return int(value), nil
	case int32:
		return int(value), nil
	case inter:
		return value.Int()
	default:
		return 0, internalCannotCastComplainer{expectedType:"int", actualType:typeof(value)}
	}
}

func MustInt(v interface{}) (int, error) {

	x, err := Int(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type inter interface {
	Int() (int, error)
}
