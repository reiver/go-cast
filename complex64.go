package cast

func Complex64(v interface{}) (complex64, error) {

	switch value := v.(type) {
	case complex64:
		return complex64(value), nil
	case float32:
		return complex(float32(value), 0), nil
	case uint8:
		return complex(float32(value), 0), nil
	case uint16:
		return complex(float32(value), 0), nil
	case int8:
		return complex(float32(value), 0), nil
	case int16:
		return complex(float32(value), 0), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"complex64", actualType:typeof(value)}
	}
}

func MustComplex64(v interface{}) (complex64, error) {

	x, err := Complex64(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}
