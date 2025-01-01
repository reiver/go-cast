package cast

// Float32 will return an float32 when `v` is of type float32, int16, int8, uint16, uint8 or has a method:
//
//	type interface {
//		Float32() (float32, error)
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
func Float32(v any) (float32, error) {

	switch value := v.(type) {
	case float32er:
		return value.Float32()
	case int16er:
		return func()(float32, error){
			casted, err := value.Int16()
			if nil != err {
				return 0, err
			}
			return float32(casted), nil
		}()
	case int8er:
		return func()(float32, error){
			casted, err := value.Int8()
			if nil != err {
				return 0, err
			}
			return float32(casted), nil
		}()
	case uint16er:
		return func()(float32, error){
			casted, err := value.Uint16()
			if nil != err {
				return 0, err
			}
			return float32(casted), nil
		}()
	case uint8er:
		return func()(float32, error){
			casted, err := value.Uint8()
			if nil != err {
				return 0, err
			}
			return float32(casted), nil
		}()
	case float32:
		return float32(value), nil
	case int16:
		return float32(value), nil
	case int8:
		return float32(value), nil
	case uint16:
		return float32(value), nil
	case uint8:
		return float32(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"float32", actualType:typeof(value)}
	}
}

// Float32Else is similar to [Float32] except that if a cast cannot be done, it returns the `alternative`.
func Float32Else(v any, alternative float32) float32 {
	result, err := Float32(v)
	if nil != err {
		return alternative
	}
	return result
}

// MustFloat32 is like Float32, expect panic()s on an error.
func MustFloat32(v any) float32 {

	x, err := Float32(v)
	if nil != err {
		panic(err)
	}

	return x
}

type float32er interface {
	Float32() (float32, error)
}
