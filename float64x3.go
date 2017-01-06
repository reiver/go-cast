package cast

func Float64x3(v interface{}) ([3]float64, error) {

	switch value := v.(type) {
	case [3]float32:
		return [3]float64{float64(value[0]), float64(value[1]), float64(value[2])}, nil
	case [3]float64:
		return [3]float64(value), nil
	default:
		return [3]float64{0.0,0.0}, internalCannotCastComplainer{expectedType:"[3]float64", actualType:typeof(value)}
	}
}

func MustFloat64x3(v interface{}) ([3]float64, error) {

	x, err := Float64x3(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}
