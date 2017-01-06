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

// MustTime is like Time, expect panic()s on an error.
func MustTime(v interface{}) time.Time {

	x, err := Time(v)
	if nil != err {
		panic(err)
	}

	return x
}

type timer interface {
	Time() (time.Time, error)
}
