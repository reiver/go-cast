package cast

import (
	"math"

	"testing"
)

func TestFloat64FromInt32er(t *testing.T) {

	tests := []struct{
		Value    int32er
		Expected int32
	}{
		{
			Value: testInt32erMin(),
			Expected:    math.MinInt32,
		},
		{
			Value: testInt32erNegativeOne(),
			Expected:                -1,
		},
		{
			Value: testInt32erZero(),
			Expected:         0,
		},
		{
			Value: testInt32erOne(),
			Expected:         1,
		},
		{
			Value: testInt32erMax(),
			Expected:    math.MaxInt32,
		},
	}


	for testNumber, test := range tests {

		x, err := Float64(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := int32(x)

		if expected, actual := test.Expected, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
