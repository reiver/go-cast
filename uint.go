package cast

// Uint will return an int when `v` is of type uint32, uint16, uint8, int or has a method:
//
//      type interface {
//              Uint32() (uint32, error)
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
func Uint(v any) (uint, error) {

	switch value := v.(type) {
	case uinter:
		return value.Uint()
	case uint32er:
		return func()(uint, error){
			casted, err := value.Uint32()
			if nil != err {
				return 0, err
			}
			return uint(casted), nil
		}()
	case uint16er:
		return func()(uint, error){
			casted, err := value.Uint16()
			if nil != err {
				return 0, err
			}
			return uint(casted), nil
		}()
	case uint8er:
		return func()(uint, error){
			casted, err := value.Uint8()
			if nil != err {
				return 0, err
			}
			return uint(casted), nil
		}()
	case uint:
		return uint(value), nil
	case uint32:
		return uint(value), nil
	case uint16:
		return uint(value), nil
	case uint8:
		return uint(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"uint", actualType:typeof(value)}
	}
}

// MustUint is like Uint, expect panic()s on an error.
func MustUint(v any) uint {

	x, err := Uint(v)
	if nil != err {
		panic(err)
	}

	return x
}

type uinter interface {
	Uint() (uint, error)
}
