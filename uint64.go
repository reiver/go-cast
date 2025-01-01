package cast

// Uint64 will return an int64 when `v` is of type uint64, uint32, uint16, uint8, uint or has a method:
//
//	type interface {
//		Uint64() (uint64, error)
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
// ... that returns successfully, or has a method:
//
//	type interface {
//		Uint() (uint, error)
//	}
//
// ... that returns successfully.
//
// Else it will return an error.
func Uint64(v any) (uint64, error) {

	switch value := v.(type) {
	case uint64er:
		return value.Uint64()
	case uint32er:
		return func()(uint64, error){
			casted, err := value.Uint32()
			if nil != err {
				return 0, err
			}
			return uint64(casted), nil
		}()
	case uint16er:
		return func()(uint64, error){
			casted, err := value.Uint16()
			if nil != err {
				return 0, err
			}
			return uint64(casted), nil
		}()
	case uint8er:
		return func()(uint64, error){
			casted, err := value.Uint8()
			if nil != err {
				return 0, err
			}
			return uint64(casted), nil
		}()
	case uinter:
		return func()(uint64, error) {
			casted, err := value.Uint()
			if nil != err {
				return 0, err
			}
			return uint64(casted), nil
		}()
	case uint64:
		return uint64(value), nil
	case uint32:
		return uint64(value), nil
	case uint16:
		return uint64(value), nil
	case uint8:
		return uint64(value), nil
	case uint:
		return uint64(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"uint64", actualType:typeof(value)}
	}
}

// MustUint64 is like Uint64, expect panic()s on an error.
func MustUint64(v any) uint64 {

	x, err := Uint64(v)
	if nil != err {
		panic(err)
	}

	return x
}

type uint64er interface {
	 Uint64() (uint64, error)
}
