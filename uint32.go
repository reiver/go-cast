package cast

func Uint32(v interface{}) (uint32, error) {

	switch value := v.(type) {
	case uint8:
		return uint32(value), nil
	case uint16:
		return uint32(value), nil
	case uint32:
		return uint32(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"uint32", actualType:typeof(value)}
	}
}

func MustUint32(v interface{}) (uint32, error) {

	x, err := Uint32(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}
