package cast

func Complex128(v any) (complex128, error) {

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

// Complex128Else is similar to [Complex128] except that if a cast cannot be done, it returns the `alternative`.
func Complex128Else(v any, alternative complex128) complex128 {
	result, err := Complex128(v)
	if nil != err {
		return alternative
	}
	return result
}

// MustComplex128 is like Complex128, expect panic()s on an error.
func MustComplex128(v any) complex128 {

	x, err := Complex128(v)
	if nil != err {
		panic(err)
	}

	return x
}

type complex128er interface {
	Complex128() (complex128, error)
}
