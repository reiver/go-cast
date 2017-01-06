package cast

func Float64x4(v interface{}) ([4]float64, error) {

	switch value := v.(type) {
	case [4]float32:
		return [4]float64{float64(value[0]), float64(value[1]), float64(value[2]), float64(value[3])}, nil
	case [4]float64:
		return [4]float64(value), nil
	default:
		return [4]float64{0.0,0.0}, internalCannotCastComplainer{expectedType:"[4]float64", actualType:typeof(value)}
	}
}

func MustFloat64x4(v interface{}) ([4]float64, error) {

	x, err := Float64x4(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}
