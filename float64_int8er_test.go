package cast

import (
	"math"

	"testing"
)

func TestFloat64FromInt8er(t *testing.T) {

	tests := []struct{
		Value    int8er
		Expected int8
	}{
		{
			Value: testInt8erMin(),
			Expected:   math.MinInt8,
		},
		{
			Value: testInt8erNegativeOne(),
			Expected:               -1,
		},
		{
			Value: testInt8erZero(),
			Expected:        0,
		},
		{
			Value: testInt8erOne(),
			Expected:        1,
		},
		{
			Value: testInt8erMax(),
			Expected:   math.MaxInt8,
		},
	}


	for testNumber, test := range tests {

		x, err := Float64(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := int8(x)

		if expected, actual := test.Expected, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
