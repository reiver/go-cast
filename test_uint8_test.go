package cast

import (
	"errors"
	"math"
)

type testUint8er struct {
	value string
}

func testUint8erZero() testUint8er {
	return testUint8er{value: "zero"}
}

func testUint8erOne() testUint8er {
	return testUint8er{value: "one"}
}

func testUint8erMax() testUint8er {
	return testUint8er{value: "max"}
}

func (receiver testUint8er) Uint8() (uint8, error) {
	switch receiver.value {
	case "zero":
		return 0, nil
	case "one":
		return 1, nil
	case "max":
		return math.MaxUint8, nil
	default:
		return 0, errors.New("Internal Error")
	}
}
