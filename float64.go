package cast

// Float64 will return an float64 when `v` is of type float64, float32, int32, int16, int8, uint32, uint16, uint8 or has a method:
//
//	type interface {
//		Float64() (float64, error)
//	}
//
// ... that returns successfully, or has a method:
//
//	type interface {
//		Float32() (float32, error)
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
func Float64(v any) (float64, error) {

	switch value := v.(type) {
	case float64er:
		return value.Float64()
	case float32er:
		return func()(float64, error){
			casted, err := value.Float32()
			if nil != err {
				return 0, err
			}
			return float64(casted), nil
		}()
	case int32er:
		return func()(float64, error){
			casted, err := value.Int32()
			if nil != err {
				return 0, err
			}
			return float64(casted), nil
		}()
	case int16er:
		return func()(float64, error){
			casted, err := value.Int16()
			if nil != err {
				return 0, err
			}
			return float64(casted), nil
		}()
	case int8er:
		return func()(float64, error){
			casted, err := value.Int8()
			if nil != err {
				return 0, err
			}
			return float64(casted), nil
		}()
	case uint32er:
		return func()(float64, error){
			casted, err := value.Uint32()
			if nil != err {
				return 0, err
			}
			return float64(casted), nil
		}()
	case uint16er:
		return func()(float64, error){
			casted, err := value.Uint16()
			if nil != err {
				return 0, err
			}
			return float64(casted), nil
		}()
	case uint8er:
		return func()(float64, error){
			casted, err := value.Uint8()
			if nil != err {
				return 0, err
			}
			return float64(casted), nil
		}()
	case float64:
		return float64(value), nil
	case float32:
		return float64(value), nil
	case int32:
		return float64(value), nil
	case int16:
		return float64(value), nil
	case int8:
		return float64(value), nil
	case uint32:
		return float64(value), nil
	case uint16:
		return float64(value), nil
	case uint8:
		return float64(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"float64", actualType:typeof(value)}
	}
}

// Float64Else is similar to [Float64] except that if a cast cannot be done, it returns the `alternative`.
func Float64Else(v any, alternative float64) float64 {
	result, err := Float64(v)
	if nil != err {
		return alternative
	}
	return result
}

// MustFloat64 is like Float64, expect panic()s on an error.
func MustFloat64(v any) float64 {

	x, err := Float64(v)
	if nil != err {
		panic(err)
	}

	return x
}

type float64er interface {
	Float64() (float64, error)
}
