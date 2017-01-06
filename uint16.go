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

// MustUint16 is like Uint16, expect panic()s on an error.
func MustUint16(v interface{}) uint16 {

	x, err := Uint16(v)
	if nil != err {
		panic(err)
	}

	return x
}

type uint16er interface {
	Uint16() (uint16, error)
}
