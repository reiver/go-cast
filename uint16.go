package cast

func Uint16(v interface{}) (uint16, error) {

	switch value := v.(type) {
	case uint8:
		return uint16(value), nil
	case uint16:
		return uint16(value), nil
	case uint16er:
		return value.Uint16()
	default:
		return 0, internalCannotCastComplainer{expectedType:"uint16", actualType:typeof(value)}
	}
}

func MustUint16(v interface{}) (uint16, error) {

	x, err := Uint16(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type uint16er interface {
	Uint16() (uint16, error)
}
