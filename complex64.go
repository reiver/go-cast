package cast

// Complex64 will return a complex64 when `v` is of type complex64, float32, uint8, uint16, int8, int16, or has a method:
//
//	type interface {
//		Complex64() (complex64, error)
//	}
//
// Else it will return an error.
//
// When float32, uint8, uint16, int8, int16, are converted to a complex64, their value goes into the "real" component
// of the complex number.
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
	case complex64er:
		return value.Complex64()
	default:
		return 0, internalCannotCastComplainer{expectedType:"complex64", actualType:typeof(value)}
	}
}

// MustComplex64 is like Complex64, expect panic()s on an error.
func MustComplex64(v interface{}) complex64 {

	x, err := Complex64(v)
	if nil != err {
		panic(err)
	}

	return x
}

type complex64er interface {
	Complex64() (complex64, error)
}
