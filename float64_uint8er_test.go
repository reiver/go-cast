package cast

import (
	"math"

	"testing"
)

func TestFloat64FromUint8er(t *testing.T) {

	tests := []struct{
		Value    uint8er
		Expected uint8
	}{
		{
			Value: testUint8erZero(),
			Expected:         0,
		},
		{
			Value: testUint8erOne(),
			Expected:         1,
		},
		{
			Value: testUint8erMax(),
			Expected:    math.MaxUint8,
		},
	}


	for testNumber, test := range tests {

		x, err := Float64(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := uint8(x)

		if expected, actual := test.Expected, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
