package cast

func String(v interface{}) (string, error) {

	switch value := v.(type) {
	case []byte:
		return string(value), nil
	case rune:
		return string(value), nil
	case []rune:
		return string(value), nil
	case string:
		return string(value), nil
	case stringer:
		return value.String()
	default:
		return "", internalCannotCastComplainer{expectedType:"string", actualType:typeof(value)}
	}
}

func MustString(v interface{}) (string, error) {

	x, err := String(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type stringer interface {
	String() (string, error)
}
