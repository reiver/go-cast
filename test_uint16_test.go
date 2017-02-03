package cast

import (
	"errors"
	"math"
)

type testUint16er struct {
	value string
}

func testUint16erZero() testUint16er {
	return testUint16er{value: "zero"}
}

func testUint16erOne() testUint16er {
	return testUint16er{value: "one"}
}

func testUint16erMax() testUint16er {
	return testUint16er{value: "max"}
}

func (receiver testUint16er) Uint16() (uint16, error) {
	switch receiver.value {
	case "zero":
		return 0, nil
	case "one":
		return 1, nil
	case "max":
		return math.MaxUint16, nil
	default:
		return 0, errors.New("Internal Error")
	}
}
