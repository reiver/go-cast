package cast

import (
	"math"

	"testing"
)

func TestFloat64FromUint32(t *testing.T) {

	tests := []struct{
		Value uint32
	}{
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: math.MaxUint32,
		},
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			test := struct{
				Value uint32
			}{
				Value: uint32(randomness.Int63n(math.MaxUint32)),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		x, err := Float64(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := uint32(x)

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
