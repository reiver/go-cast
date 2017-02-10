package cast

import (
	"math"

	"testing"
)

func TestFloat64FromInt16er(t *testing.T) {

	tests := []struct{
		Value    int16er
		Expected int16
	}{
		{
			Value: testInt16erMin(),
			Expected:    math.MinInt16,
		},
		{
			Value: testInt16erNegativeOne(),
			Expected:                -1,
		},
		{
			Value: testInt16erZero(),
			Expected:         0,
		},
		{
			Value: testInt16erOne(),
			Expected:         1,
		},
		{
			Value: testInt16erMax(),
			Expected:    math.MaxInt16,
		},
	}


	for testNumber, test := range tests {

		x, err := Float64(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := int16(x)

		if expected, actual := test.Expected, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
