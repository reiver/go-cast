package cast

func Uint32(v interface{}) (uint32, error) {

	switch value := v.(type) {
	case uint8:
		return uint32(value), nil
	case uint16:
		return uint32(value), nil
	case uint32:
		return uint32(value), nil
	case uint32er:
		return value.Uint32()
	default:
		return 0, internalCannotCastComplainer{expectedType:"uint32", actualType:typeof(value)}
	}
}

// MustUint32 is like Uint32, expect panic()s on an error.
func MustUint32(v interface{}) uint32 {

	x, err := Uint32(v)
	if nil != err {
		panic(err)
	}

	return x
}

type uint32er interface {
	Uint32() (uint32, error)
}
