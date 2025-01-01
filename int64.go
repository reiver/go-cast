package cast

// Int64 will return an int64 when `v` is of type int64, int32, int16, int8, int, uint32, uint16, uint8 or has a method:
//
//	type interface {
//		Int64() (int64, error)
//	}
//
// ... that returns successfully, or has a method:
//
//	type interface {
//		Int32() (int32, error)
//	}
//
// ... that returns successfully, or has a method:
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
//		Int() (int, error)
//	}
//
// ... that returns successfully, or has a method:
//
//	type interface {
//		Uint32() (uint32, error)
//	}
//
// ... that returns successfully, or has a method:
//
//	type interface {
//		Uint16() (uint16, error)
//	}
//
// ... that returns successfully, or has a method:
//
//	type interface {
//		Uint8() (uint8, error)
//	}
//
// ... that returns successfully.
//
// Else it will return an error.
func Int64(v any) (int64, error) {

	switch value := v.(type) {
	case int64er:
		return value.Int64()
	case int32er:
		return func()(int64, error){
			casted, err := value.Int32()
			if nil != err {
				return 0, err
			}
			return int64(casted), nil
		}()
	case int16er:
		return func()(int64, error){
			casted, err := value.Int16()
			if nil != err {
				return 0, err
			}
			return int64(casted), nil
		}()
	case int8er:
		return func()(int64, error){
			casted, err := value.Int8()
			if nil != err {
				return 0, err
			}
			return int64(casted), nil
		}()
	case inter:
		return func()(int64, error){
			casted, err := value.Int()
			if nil != err {
				return 0, err
			}
			return int64(casted), nil
		}()
	case uint32er:
		return func()(int64, error){
			casted, err := value.Uint32()
			if nil != err {
				return 0, err
			}
			return int64(casted), nil
		}()
	case uint16er:
		return func()(int64, error){
			casted, err := value.Uint16()
			if nil != err {
				return 0, err
			}
			return int64(casted), nil
		}()
	case uint8er:
		return func()(int64, error){
			casted, err := value.Uint8()
			if nil != err {
				return 0, err
			}
			return int64(casted), nil
		}()
	case int64:
		return int64(value), nil
	case int32:
		return int64(value), nil
	case int16:
		return int64(value), nil
	case int8:
		return int64(value), nil
	case int:
		return int64(value), nil
	case uint32:
		return int64(value), nil
	case uint16:
		return int64(value), nil
	case uint8:
		return int64(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"int64", actualType:typeof(value)}
	}
}

// Int64Else is similar to [Int64] except that if a cast cannot be done, it returns the `alternative`.
func Int64Else(v any, alternative int64) int64 {
	result, err := Int64(v)
	if nil != err {
		return alternative
	}
	return result
}

// MustInt64 is like Int64, expect panic()s on an error.
func MustInt64(v any) int64 {

	x, err := Int64(v)
	if nil != err {
		panic(err)
	}

	return x
}

type int64er interface {
	Int64() (int64, error)
}
