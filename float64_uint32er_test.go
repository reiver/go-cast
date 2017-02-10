package cast

import (
	"math"

	"testing"
)

func TestFloat64FromUint32er(t *testing.T) {

	tests := []struct{
		Value    uint32er
		Expected uint32
	}{
		{
			Value: testUint32erZero(),
			Expected:          0,
		},
		{
			Value: testUint32erOne(),
			Expected:          1,
		},
		{
			Value: testUint32erMax(),
			Expected:     math.MaxUint32,
		},
	}


	for testNumber, test := range tests {

		x, err := Float64(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := uint32(x)

		if expected, actual := test.Expected, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
