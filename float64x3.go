package cast

func Float64x3(v interface{}) ([3]float64, error) {

	switch value := v.(type) {
	case [3]float32:
		return [3]float64{float64(value[0]), float64(value[1]), float64(value[2])}, nil
	case [3]float64:
		return [3]float64(value), nil
	case float64x3er:
		return value.Float64x3()
	default:
		return [3]float64{0.0,0.0}, internalCannotCastComplainer{expectedType:"[3]float64", actualType:typeof(value)}
	}
}

// MustFloat64x3 is like Float64x3, expect panic()s on an error.
func MustFloat64x3(v interface{}) [3]float64 {

	x, err := Float64x3(v)
	if nil != err {
		panic(err)
	}

	return x
}

type float64x3er interface {
	Float64x3() ([3]float64, error)
}
