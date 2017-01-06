package cast

import (
	"time"
)

func Time(v interface{}) (time.Time, error) {

	switch value := v.(type) {
	case time.Time:
		return time.Time(value), nil
	case timer:
		return value.Time()
	default:
		return time.Time{}, internalCannotCastComplainer{expectedType:"time.Time", actualType:typeof(value)}
	}
}

func MustTime(v interface{}) (time.Time, error) {

	x, err := Time(v)
	if nil != err {
		panic(err)
	}

	return x, nil
}

type timer interface {
	Time() (time.Time, error)
}
