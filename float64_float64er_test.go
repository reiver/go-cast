package cast

import (
	"math"

	"testing"
)

func TestFloat64FromFloat64er(t *testing.T) {

	tests := []struct{
		Value    float64er
		Expected float64
	}{
		{
			Value: testFloat64erNegativeInfinity(),
			Expected:              math.Inf(-1),
		},
		{
			Value: testFloat64erMin(),
			Expected:     -math.MaxFloat64,
		},
		{
			Value: testFloat64erNegativePi(),
			Expected:             -math.Pi,
		},
		{
			Value: testFloat64erNegativeE(),
			Expected:             -math.E,
		},
		{
			Value: testFloat64erNegativeSqrt2(),
			Expected:             -math.Sqrt2,
		},
		{
			Value: testFloat64erNegativeOne(),
			Expected:                  -1.0,
		},
		{
			Value: testFloat64erNegativeLn2(),
			Expected:             -math.Ln2,
		},
		{
			Value: testFloat64erNegativeSmallestNonzero(),
			Expected:             -math.SmallestNonzeroFloat64,
		},
		{
			Value: testFloat64erZero(),
			Expected:           0.0,
		},
		{
			Value: testFloat64erSmallestNonzero(),
			Expected:      math.SmallestNonzeroFloat64,
		},
		{
			Value: testFloat64erLn2(),
			Expected:       math.Ln2,
		},
		{
			Value: testFloat64erOne(),
			Expected:           1.0,
		},
		{
			Value: testFloat64erSqrt2(),
			Expected:      math.Sqrt2,
		},
		{
			Value: testFloat64erE(),
			Expected:      math.E,
		},
		{
			Value: testFloat64erPi(),
			Expected:      math.Pi,
		},
		{
			Value: testFloat64erMax(),
			Expected:      math.MaxFloat64,
		},
		{
			Value: testFloat64erInfinity(),
			Expected:      math.Inf(+1),
		},



		{
			Value: testFloat64erNaN(),
			Expected:      math.NaN(),
		},
	}


	for testNumber, test := range tests {

		x, err := Float64(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := float64(x)

		if expected, actual := test.Expected, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
