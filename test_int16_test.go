package cast

import (
	"errors"
	"math"
)

type testInt16er struct {
	value string
}

func testInt16erMin() testInt16er {
	return testInt16er{value: "min"}
}

func testInt16erNegativeOne() testInt16er {
	return testInt16er{value: "negative one"}
}

func testInt16erZero() testInt16er {
	return testInt16er{value: "zero"}
}

func testInt16erOne() testInt16er {
	return testInt16er{value: "one"}
}

func testInt16erMax() testInt16er {
	return testInt16er{value: "max"}
}

func (receiver testInt16er) Int16() (int16, error) {
	switch receiver.value {
	case "min":
		return math.MinInt16, nil
	case "negative one":
		return -1, nil
	case "zero":
		return 0, nil
	case "one":
		return 1, nil
	case "max":
		return math.MaxInt16, nil
	default:
		return 0, errors.New("Internal Error")
	}
}
