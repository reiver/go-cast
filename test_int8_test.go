package cast

import (
	"errors"
	"math"
)

type testInt8er struct {
	value string
}

func testInt8erMin() testInt8er {
	return testInt8er{value: "min"}
}

func testInt8erNegativeOne() testInt8er {
	return testInt8er{value: "negative one"}
}

func testInt8erZero() testInt8er {
	return testInt8er{value: "zero"}
}

func testInt8erOne() testInt8er {
	return testInt8er{value: "one"}
}

func testInt8erMax() testInt8er {
	return testInt8er{value: "max"}
}

func (receiver testInt8er) Int8() (int8, error) {
	switch receiver.value {
	case "min":
		return math.MinInt8, nil
	case "negative one":
		return -1, nil
	case "zero":
		return 0, nil
	case "one":
		return 1, nil
	case "max":
		return math.MaxInt8, nil
	default:
		return 0, errors.New("Internal Error")
	}
}
