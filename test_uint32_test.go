package cast

import (
	"errors"
	"math"
)

type testUint32er struct {
	value string
}

func testUint32erZero() testUint32er {
	return testUint32er{value: "zero"}
}

func testUint32erOne() testUint32er {
	return testUint32er{value: "one"}
}

func testUint32erMax() testUint32er {
	return testUint32er{value: "max"}
}

func (receiver testUint32er) Uint32() (uint32, error) {
	switch receiver.value {
	case "zero":
		return 0, nil
	case "one":
		return 1, nil
	case "max":
		return math.MaxUint32, nil
	default:
		return 0, errors.New("Internal Error")
	}
}
