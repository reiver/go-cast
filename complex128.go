package cast

func Complex128(v interface{}) (complex128, error) {

	switch value := v.(type) {
	case complex64:
		return complex128(value), nil
	case complex128:
		return complex128(value), nil
	case float32:
		return complex(float64(value), 0), nil
	case float64:
		return complex(float64(value), 0), nil
	case uint8:
		return complex(float64(value), 0), nil
	case uint16:
		return complex(float64(value), 0), nil
	case int8:
		return complex(float64(value), 0), nil
	case int16:
		return complex(float64(value), 0), nil
	case complex128er:
		return value.Complex128()
	default:
		return 0, internalCannotCastComplainer{expectedType:"complex128", actualType:typeof(value)}
	}
}

func MustComplex128(v interface{}) (complex128, error) {

	x, err := Complex128(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type complex128er interface {
	Complex128() (complex128, error)
}
