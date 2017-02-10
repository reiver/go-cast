package cast

import (
	"math"

	"testing"
)

func TestFloat64FromFloat32er(t *testing.T) {

	tests := []struct{
		Value    float32er
		Expected float32
	}{
		{
			Value: testFloat32erNegativeInfinity(),
			Expected:      float32(math.Inf(-1)),
		},
		{
			Value: testFloat32erMin(),
			Expected:     -math.MaxFloat32,
		},
		{
			Value: testFloat32erNegativePi(),
			Expected:             -math.Pi,
		},
		{
			Value: testFloat32erNegativeE(),
			Expected:             -math.E,
		},
		{
			Value: testFloat32erNegativeSqrt2(),
			Expected:             -math.Sqrt2,
		},
		{
			Value: testFloat32erNegativeOne(),
			Expected:                  -1.0,
		},
		{
			Value: testFloat32erNegativeLn2(),
			Expected:             -math.Ln2,
		},
		{
			Value: testFloat32erNegativeSmallestNonzero(),
			Expected:             -math.SmallestNonzeroFloat32,
		},
		{
			Value: testFloat32erZero(),
			Expected:           0.0,
		},
		{
			Value: testFloat32erSmallestNonzero(),
			Expected:      math.SmallestNonzeroFloat32,
		},
		{
			Value: testFloat32erLn2(),
			Expected:      math.Ln2,
		},
		{
			Value: testFloat32erOne(),
			Expected:           1.0,
		},
		{
			Value: testFloat32erSqrt2(),
			Expected:      math.Sqrt2,
		},
		{
			Value: testFloat32erE(),
			Expected:      math.E,
		},
		{
			Value: testFloat32erPi(),
			Expected:      math.Pi,
		},
		{
			Value: testFloat32erMax(),
			Expected:      math.MaxFloat32,
		},
		{
			Value:    testFloat32erInfinity(),
			Expected: float32(math.Inf(+1)),
		},



		{
			Value:    testFloat32erNaN(),
			Expected: float32(math.NaN()),
		},
	}


	for testNumber, test := range tests {

		x, err := Float64(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := float32(x)

		if expected, actual := test.Expected, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
