package cast

import (
	"time"
)

func Time(v any) (time.Time, error) {

	switch value := v.(type) {
	case timer:
		return value.Time()
	case time.Time:
		return time.Time(value), nil
	default:
		return time.Time{}, internalCannotCastComplainer{expectedType:"time.Time", actualType:typeof(value)}
	}
}

// MustTime is like Time, expect panic()s on an error.
func MustTime(v any) time.Time {

	x, err := Time(v)
	if nil != err {
		panic(err)
	}

	return x
}

type timer interface {
	Time() (time.Time, error)
}
