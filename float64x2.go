package cast

func Float64x2(v interface{}) ([2]float64, error) {

	switch value := v.(type) {
	case [2]float32:
		return [2]float64{float64(value[0]), float64(value[1])}, nil
	case [2]float64:
		return [2]float64(value), nil
	default:
		return [2]float64{0.0,0.0}, internalCannotCastComplainer{expectedType:"[2]float64", actualType:typeof(value)}
	}
}

func MustFloat64x2(v interface{}) ([2]float64, error) {

	x, err := Float64x2(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}
