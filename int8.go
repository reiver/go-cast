package cast

func Int8(v interface{}) (int8, error) {

	switch value := v.(type) {
	case int8:
		return int8(value), nil
	case int8er:
		return value.Int8()
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

type int8er interface {
	Int8() (int8, error)
}
