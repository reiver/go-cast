package cast

import (
	"errors"
	"math"
)

type testInt32er struct {
	value string
}

func testInt32erMin() testInt32er {
	return testInt32er{value: "min"}
}

func testInt32erNegativeOne() testInt32er {
	return testInt32er{value: "negative one"}
}

func testInt32erZero() testInt32er {
	return testInt32er{value: "zero"}
}

func testInt32erOne() testInt32er {
	return testInt32er{value: "one"}
}

func testInt32erMax() testInt32er {
	return testInt32er{value: "max"}
}

func (receiver testInt32er) Int32() (int32, error) {
	switch receiver.value {
	case "min":
		return math.MinInt32, nil
	case "negative one":
		return -1, nil
	case "zero":
		return 0, nil
	case "one":
		return 1, nil
	case "max":
		return math.MaxInt32, nil
	default:
		return 0, errors.New("Internal Error")
	}
}
