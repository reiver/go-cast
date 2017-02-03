package cast

import (
	"errors"
)

type testBooler struct {
	value string
}

func testBoolerTrue() testBooler {
	return testBooler{value: "true"}
}

func testBoolerFalse() testBooler {
	return testBooler{value: "false"}
}

func (receiver testBooler) Bool() (bool, error) {
	switch receiver.value {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, errors.New("Internal Error")
	}
}
