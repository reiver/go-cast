package cast

// Int8 will return a int8 when `v` is of type int8, or has a method:
//
//	type interface {
//		Int8() (bool, error)
//	}
//
// ... that returns successfully.
//
// Else it will return an error.
func Int8(v interface{}) (int8, error) {

	switch value := v.(type) {
	case int8:
		return int8(value), nil
	case int8er:
		return value.Int8()
	default:
		return 0, internalCannotCastComplainer{expectedType:"int8", actualType:typeof(value)}
	}
}

// MustInt8 is like Int8, expect panic()s on an error.
func MustInt8(v interface{}) int8 {

	x, err := Int8(v)
	if nil != err {
		panic(err)
	}

	return x
}

type int8er interface {
	Int8() (int8, error)
}
