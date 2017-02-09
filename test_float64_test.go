package cast

import (
	"errors"
	"math"
)

type testFloat64er struct {
	value string
}

func testFloat64erNegativeInfinity() testFloat64er {
	return testFloat64er{value: "negative infinity"}
}

func testFloat64erMin() testFloat64er {
	return testFloat64er{value: "min"}
}

func testFloat64erNegativePi() testFloat64er {
	return testFloat64er{value: "negative pi"}
}

func testFloat64erNegativeE() testFloat64er {
	return testFloat64er{value: "negative e"}
}

func testFloat64erNegativeSqrt2() testFloat64er {
	return testFloat64er{value: "negative sqrt(2)"}
}

func testFloat64erNegativeOne() testFloat64er {
	return testFloat64er{value: "negative one"}
}

func testFloat64erNegativeLn2() testFloat64er {
	return testFloat64er{value: "negative ln(2)"}
}

func testFloat64erNegativeSmallestNonzero() testFloat64er {
	return testFloat64er{value: "negative smallest-nonzero"}
}

func testFloat64erZero() testFloat64er {
	return testFloat64er{value: "zero"}
}

func testFloat64erSmallestNonzero() testFloat64er {
	return testFloat64er{value: "smallest-nonzero"}
}

func testFloat64erLn2() testFloat64er {
	return testFloat64er{value: "ln(2)"}
}

func testFloat64erOne() testFloat64er {
	return testFloat64er{value: "one"}
}

func testFloat64erSqrt2() testFloat64er {
	return testFloat64er{value: "sqrt(2)"}
}

func testFloat64erE() testFloat64er {
	return testFloat64er{value: "e"}
}

func testFloat64erPi() testFloat64er {
	return testFloat64er{value: "pi"}
}

func testFloat64erMax() testFloat64er {
	return testFloat64er{value: "max"}
}

func testFloat64erInfinity() testFloat64er {
	return testFloat64er{value: "infinity"}
}

func testFloat64erNaN() testFloat64er {
	return testFloat64er{value: "nan"}
}

func (receiver testFloat64er) Float64() (float64, error) {
	switch receiver.value {
	case "negative infinity":
		return math.Inf(-1), nil
	case "min":
		return -math.MaxFloat64, nil
	case "negative pi":
		return -math.Pi, nil
	case "negative e":
		return -math.E, nil
	case "negative sqrt(2)":
		return -math.Sqrt2, nil
	case "negative one":
		return -1, nil
	case "negative ln(2)":
		return -math.Ln2, nil
	case "negative smallest-nonzero":
		return -math.SmallestNonzeroFloat64, nil
	case "zero":
		return 0, nil
	case "smallest-nonzero":
		return math.SmallestNonzeroFloat64, nil
	case "ln(2)":
		return math.Ln2, nil
	case "one":
		return 1, nil
	case "sqrt(2)":
		return math.Sqrt2, nil
	case "e":
		return math.E, nil
	case "pi":
		return math.Pi, nil
	case "max":
		return math.MaxFloat64, nil
	case "infinity":
		return math.Inf(+1), nil
	case "nan":
		return math.NaN(), nil
	default:
		return 0, errors.New("Internal Error")
	}
}
