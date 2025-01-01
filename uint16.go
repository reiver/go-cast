package cast

// Uint16 will return an int16 when `v` is of type uint16, uint8 or has a method:
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
// ... that returns successfully.
//
// Else it will return an error.
func Uint16(v any) (uint16, error) {

	switch value := v.(type) {
	case uint16er:
		return value.Uint16()
	case uint8er:
		return func()(uint16, error){
			casted, err := value.Uint8()
			if nil != err {
				return 0, err
			}
			return uint16(casted), nil
		}()
	case uint16:
		return uint16(value), nil
	case uint8:
		return uint16(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"uint16", actualType:typeof(value)}
	}
}

// MustUint16 is like Uint16, expect panic()s on an error.
func MustUint16(v any) uint16 {

	x, err := Uint16(v)
	if nil != err {
		panic(err)
	}

	return x
}

type uint16er interface {
	Uint16() (uint16, error)
}
