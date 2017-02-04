package cast

import (
	"errors"
	"math"
)

type testUint64er struct {
	value string
}

func testUint64erZero() testUint64er {
	return testUint64er{value: "zero"}
}

func testUint64erOne() testUint64er {
	return testUint64er{value: "one"}
}

func testUint64erMax() testUint64er {
	return testUint64er{value: "max"}
}

func (receiver testUint64er) Uint64() (uint64, error) {
	switch receiver.value {
	case "zero":
		return 0, nil
	case "one":
		return 1, nil
	case "max":
		return math.MaxUint64, nil
	default:
		return 0, errors.New("Internal Error")
	}
}
