package cast

// Uint32 will return an int32 when `v` is of type uint32, uint16, uint8 or has a method:
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
func Uint32(v any) (uint32, error) {

	switch value := v.(type) {
	case uint32er:
		return value.Uint32()
	case uint16er:
		return func()(uint32, error){
			casted, err := value.Uint16()
			if nil != err {
				return 0, err
			}
			return uint32(casted), nil
		}()
	case uint8er:
		return func()(uint32, error){
			casted, err := value.Uint8()
			if nil != err {
				return 0, err
			}
			return uint32(casted), nil
		}()
	case uint32:
		return uint32(value), nil
	case uint16:
		return uint32(value), nil
	case uint8:
		return uint32(value), nil
	default:
		return 0, internalCannotCastComplainer{expectedType:"uint32", actualType:typeof(value)}
	}
}

// Uint32Else is similar to [Uint32] except that if a cast cannot be done, it returns the `alternative`.
func Uint32Else(v any, alternative uint32) uint32 {
	result, err := Uint32(v)
	if nil != err {
		return alternative
	}
	return result
}

// MustUint32 is like Uint32, expect panic()s on an error.
func MustUint32(v any) uint32 {

	x, err := Uint32(v)
	if nil != err {
		panic(err)
	}

	return x
}

type uint32er interface {
	Uint32() (uint32, error)

}
