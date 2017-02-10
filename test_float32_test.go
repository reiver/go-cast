package cast

import (
	"errors"
	"math"
)

type testFloat32er struct {
	value string
}

func testFloat32erNegativeInfinity() testFloat32er {
	return testFloat32er{value: "negative infinity"}
}

func testFloat32erMin() testFloat32er {
	return testFloat32er{value: "min"}
}

func testFloat32erNegativePi() testFloat32er {
	return testFloat32er{value: "negative pi"}
}

func testFloat32erNegativeE() testFloat32er {
	return testFloat32er{value: "negative e"}
}

func testFloat32erNegativeSqrt2() testFloat32er {
	return testFloat32er{value: "negative sqrt(2)"}
}

func testFloat32erNegativeOne() testFloat32er {
	return testFloat32er{value: "negative one"}
}

func testFloat32erNegativeLn2() testFloat32er {
	return testFloat32er{value: "negative ln(2)"}
}

func testFloat32erNegativeSmallestNonzero() testFloat32er {
	return testFloat32er{value: "negative smallest-nonzero"}
}

func testFloat32erZero() testFloat32er {
	return testFloat32er{value: "zero"}
}

func testFloat32erSmallestNonzero() testFloat32er {
	return testFloat32er{value: "smallest-nonzero"}
}

func testFloat32erLn2() testFloat32er {
	return testFloat32er{value: "ln(2)"}
}

func testFloat32erOne() testFloat32er {
	return testFloat32er{value: "one"}
}

func testFloat32erSqrt2() testFloat32er {
	return testFloat32er{value: "sqrt(2)"}
}

func testFloat32erE() testFloat32er {
	return testFloat32er{value: "e"}
}

func testFloat32erPi() testFloat32er {
	return testFloat32er{value: "pi"}
}

func testFloat32erMax() testFloat32er {
	return testFloat32er{value: "max"}
}

func testFloat32erInfinity() testFloat32er {
	return testFloat32er{value: "infinity"}
}

func testFloat32erNaN() testFloat32er {
	return testFloat32er{value: "nan"}
}

func (receiver testFloat32er) Float32() (float32, error) {
	switch receiver.value {
	case "negative infinity":
		return float32(math.Inf(-1)), nil
	case "min":
		return -math.MaxFloat32, nil
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
		return -math.SmallestNonzeroFloat32, nil
	case "zero":
		return 0, nil
	case "smallest-nonzero":
		return math.SmallestNonzeroFloat32, nil
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
		return math.MaxFloat32, nil
	case "infinity":
		return float32(math.Inf(+1)), nil
	case "nan":
		return float32(math.NaN()), nil
	default:
		return 0, errors.New("Internal Error")
	}
}
