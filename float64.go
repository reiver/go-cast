package cast

func Float64(v interface{}) (float64, error) {

	switch value := v.(type) {
	case float64er:
		return value.Float64()
	case float32:
		return float64(value), nil
	case float64:
		return float64(value), nil
	case uint8:
		return float64(value), nil
	case uint16:
		return float64(value), nil
	case uint32:
		return float64(value), nil
	case int8:
		return float64(value), nil
	case int16:
		return float64(value), nil
	case int32:
		return float64(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"float64", actualType:typeof(value)}
	}
}

// MustFloat64 is like Float64, expect panic()s on an error.
func MustFloat64(v interface{}) float64 {

	x, err := Float64(v)
	if nil != err {
		panic(err)
	}

	return x
}

type float64er interface {
	Float64() (float64, error)
}
