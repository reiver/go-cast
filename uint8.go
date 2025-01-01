package cast

// Uint8 will return an uint8 when `v` is of type uint8, or has a method:
//
//	type interface {
//		Uint8() (uint8, error)
//	}
//
// ... that returns successfully.
//
// Else it will return an error.
func Uint8(v any) (uint8, error) {

	switch value := v.(type) {
	case uint8er:
		return value.Uint8()
	case uint8:
		return uint8(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"uint8", actualType:typeof(value)}
	}
}

// Uint8Else is similar to [Uint8] except that if a cast cannot be done, it returns the `alternative`.
func Uint8Else(v any, alternative uint8) uint8 {
	result, err := Uint8(v)
	if nil != err {
		return alternative
	}
	return result
}

// MustUint8 is like Uint8, expect panic()s on an error.
func MustUint8(v any) uint8 {

	x, err := Uint8(v)
	if nil != err {
		panic(err)
	}

	return x
}

type uint8er interface {
	Uint8() (uint8, error)
}
