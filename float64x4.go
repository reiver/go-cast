package cast

func Float64x4(v interface{}) ([4]float64, error) {

	switch value := v.(type) {
	case [4]float32:
		return [4]float64{float64(value[0]), float64(value[1]), float64(value[2]), float64(value[3])}, nil
	case [4]float64:
		return [4]float64(value), nil
	case float64x4er:
		return value.Float64x4()
	default:
		return [4]float64{0.0,0.0}, internalCannotCastComplainer{expectedType:"[4]float64", actualType:typeof(value)}
	}
}

// MustFloat64x4 is like Float64x4, expect panic()s on an error.
func MustFloat64x4(v interface{}) [4]float64 {

	x, err := Float64x4(v)
	if nil != err {
		panic(err)
	}

	return x
}

type float64x4er interface {
	Float64x4() ([4]float64, error)
}
