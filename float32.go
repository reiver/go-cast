package cast

func Float32(v interface{}) (float32, error) {

	switch value := v.(type) {
	case float32:
		return float32(value), nil
	case uint8:
		return float32(value), nil
	case uint16:
		return float32(value), nil
	case int8:
		return float32(value), nil
	case int16:
		return float32(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"float32", actualType:typeof(value)}
	}
}

func MustFloat32(v interface{}) (float32, error) {

	x, err := Float32(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}
