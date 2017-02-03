package cast

// Int16 will return an int16 when `v` is of type int16, int8, uint8 or has a method:
//
//	type interface {
//		Int16() (int16, error)
//	}
//
// ... that returns successfully, or has a method:
//
//	type interface {
//		Int8() (int8, error)
//	}
//
// ... that returns successfully, or has a method:
//
//	type interface {
//		Uint8() (uint8, error)
//	}
//
// Else it will return an error.
func Int16(v interface{}) (int16, error) {

	switch value := v.(type) {
	case int16:
		return int16(value), nil
	case int16er:
		return value.Int16()
	case int8:
		return int16(value), nil
	case int8er:
		return func()(int16, error){
			value, err := value.Int8()
			if nil != err {
				return 0, err
			}
			return int16(value), nil
		}()
	case uint8:
		return int16(value), nil
	case uint8er:
		return func()(int16, error){
			value, err := value.Uint8()
			if nil != err {
				return 0, err
			}
			return int16(value), nil
		}()
	default:
		return 0, internalCannotCastComplainer{expectedType:"int16", actualType:typeof(value)}
	}
}

// MustInt16 is like Int16, expect panic()s on an error.
func MustInt16(v interface{}) int16 {

	x, err := Int16(v)
	if nil != err {
		panic(err)
	}

	return x
}

type int16er interface {
	Int16() (int16, error)
}
