package cast

// Int will return an int when `v` is of type int32, int16, int8, int, uint16, uint8 or has a method:
//
//      type interface {
//              Int32() (int32, error)
//      }
//
// ... that returns successfully, or has a method:
//
//      type interface {
//              Int16() (int16, error)
//      }
//
// ... that returns successfully, or has a method:
//
//      type interface {
//              Int8() (int8, error)
//      }
//
// ... that returns successfully, or has a method:
//
//      type interface {
//              Uint16() (uint16, error)
//      }
//
// ... that returns successfully, or has a method:
//
//      type interface {
//              Uint8() (uint8, error)
//      }
//
// ... that returns successfully, or has a method:
//
//      type interface {
//              Uint() (uint, error)
//      }
//
// ... that returns successfully.
//
// Else it will return an error.
func Int(v any) (int, error) {

	switch value := v.(type) {
	case inter:
		return value.Int()
	case int32er:
		return func()(int, error){
			casted, err := value.Int32()
			if nil != err {
				return 0, err
			}
			return int(casted), nil
		}()
	case int16er:
		return func()(int, error){
			casted, err := value.Int16()
			if nil != err {
				return 0, err
			}
			return int(casted), nil
		}()
	case int8er:
		return func()(int, error){
			casted, err := value.Int8()
			if nil != err {
				return 0, err
			}
			return int(casted), nil
		}()
	case uint16er:
		return func()(int, error){
			casted, err := value.Uint16()
			if nil != err {
				return 0, err
			}
			return int(casted), nil
		}()
	case uint8er:
		return func()(int, error){
			casted, err := value.Uint8()
			if nil != err {
				return 0, err
			}
			return int(casted), nil
		}()
	case int:
		return int(value), nil
	case int32:
		return int(value), nil
	case int16:
		return int(value), nil
	case int8:
		return int(value), nil
	case uint16:
		return int(value), nil
	case uint8:
		return int(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"int", actualType:typeof(value)}
	}
}

// IntElse is similar to [Int] except that if a cast cannot be done, it returns the `alternative`.
func IntElse(v any, alternative int) int {
	result, err := Int(v)
	if nil != err {
		return alternative
	}
	return result
}

// MustInt is like Int, expect panic()s on an error.
func MustInt(v any) int {

	x, err := Int(v)
	if nil != err {
		panic(err)
	}

	return x
}

type inter interface {
	Int() (int, error)
}
