package cast

func Uint8(v interface{}) (uint8, error) {

	switch value := v.(type) {
	case uint8:
		return uint8(value), nil
	case uint8er:
		return value.Uint8()
	default:
		return 0, internalCannotCastComplainer{expectedType:"uint8", actualType:typeof(value)}
	}
}

func MustUint8(v interface{}) (uint8, error) {

	x, err := Uint8(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type uint8er interface {
	Uint8() (uint8, error)
}
