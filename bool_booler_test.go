package cast

import (
	"testing"
)

func TestBoolFromBooler(t *testing.T) {

	tests := []struct{
		Value    booler
		Expected bool
	}{
		{
			Value: testBoolerFalse(),
			Expected:        false,
		},
		{
			Value: testBoolerTrue(),
			Expected:        true,
		},
	}


	for testNumber, test := range tests {

		x, err := Bool(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := bool(x)

		if expected, actual := test.Expected, y; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}
