package cast

import (
	"errors"
)

type testUinter struct {
	value string
}

func testUinterZero() testUinter {
	return testUinter{value: "zero"}
}

func testUinterOne() testUinter {
	return testUinter{value: "one"}
}

func testUinterMax() testUinter {
	return testUinter{value: "max"}
}

func (receiver testUinter) Uint() (uint, error) {
	const maxUint = ^uint(0)

	switch receiver.value {
	case "zero":
		return 0, nil
	case "one":
		return 1, nil
	case "max":
		return maxUint, nil
	default:
		return 0, errors.New("Internal Error")
	}
}
