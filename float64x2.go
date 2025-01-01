package cast

func Float64x2(v any) ([2]float64, error) {

	switch value := v.(type) {
	case float64x2er:
		return value.Float64x2()
	case [2]float64:
		return [2]float64(value), nil
	case [2]float32:
		return [2]float64{float64(value[0]), float64(value[1])}, nil
	default:
		return [2]float64{0.0,0.0}, internalCannotCastComplainer{expectedType:"[2]float64", actualType:typeof(value)}
	}
}

// Float64x2 is like Float64x2, expect panic()s on an error.
func MustFloat64x2(v any) [2]float64 {

	x, err := Float64x2(v)
	if nil != err {
		panic(err)
	}

	return x
}

type float64x2er interface {
	Float64x2() ([2]float64, error)
}
