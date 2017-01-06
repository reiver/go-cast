package cast

func Float64(v interface{}) (float64, error) {

	switch value := v.(type) {
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
	case float64er:
		return value.Float64()
	default:
		return 0, internalCannotCastComplainer{expectedType:"float64", actualType:typeof(value)}
	}
}

func MustFloat64(v interface{}) (float64, error) {

	x, err := Float64(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type float64er interface {
	Float64() (float64, error)
}
