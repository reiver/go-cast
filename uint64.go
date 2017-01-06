package cast

func Uint64(v interface{}) (uint64, error) {

	switch value := v.(type) {
	case uint:
		return uint64(value), nil
	case uint8:
		return uint64(value), nil
	case uint16:
		return uint64(value), nil
	case uint32:
		return uint64(value), nil
	case uint64:
		return uint64(value), nil
	case uint64er:
		return value.Uint64()
	default:
		return 0, internalCannotCastComplainer{expectedType:"uint64", actualType:typeof(value)}
	}
}

func MustUint64(v interface{}) (uint64, error) {

	x, err := Uint64(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type uint64er interface {
	 Uint64() (uint64, error)
}
