package cast

import (
	"math"

	"testing"
)

func TestFloat32FromFloat32(t *testing.T) {

	tests := []struct{
		Value float32
	}{
		{
			Value: float32(math.Inf(-1)),
		},
		{
			Value: -math.MaxFloat32,
		},
		{
			Value: -math.Pi,
		},
		{
			Value: -math.E,
		},
		{
			Value: -math.Sqrt2,
		},
		{
			Value: -1.0,
		},
		{
			Value: -math.Ln2,
		},
		{
			Value: -math.SmallestNonzeroFloat32,
		},
		{
			Value: 0.0,
		},
		{
			Value: math.SmallestNonzeroFloat32,
		},
		{
			Value: math.Ln2,
		},
		{
			Value: 1.0,
		},
		{
			Value: math.Sqrt2,
		},
		{
			Value: math.E,
		},
		{
			Value: math.Pi,
		},
		{
			Value: math.MaxFloat32,
		},
		{
			Value: float32(math.Inf(+1)),
		},



		{
			Value: float32(math.NaN()),
		},
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			test := struct{
				Value float32
			}{
				Value: randomness.Float32(),
			}
			tests = append(tests, test)

			test = struct{
				Value float32
			}{
				Value: -randomness.Float32(),
			}
			tests = append(tests, test)



			test = struct{
				Value float32
			}{
				Value: randomness.Float32() * math.MaxFloat32,
			}
			tests = append(tests, test)

			test = struct{
				Value float32
			}{
				Value: -randomness.Float32() * math.MaxFloat32,
			}
			tests = append(tests, test)



			test = struct{
				Value float32
			}{
				Value: randomness.Float32() * 999999999,
			}
			tests = append(tests, test)

			test = struct{
				Value float32
			}{
				Value: -randomness.Float32() * 999999999,
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		x, err := Float32(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := float32(x)

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestFloat32FromInt8(t *testing.T) {

	tests := []struct{
		Value int8
	}{
		{
			Value: math.MinInt8,
		},
		{
			Value: -1,
		},
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: math.MaxInt8,
		},
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			test := struct{
				Value int8
			}{
				Value: int8(randomness.Int63n(math.MaxInt8)),
			}
			tests = append(tests, test)

			test = struct{
				Value int8
			}{
				Value: -int8(randomness.Int63n(-1*math.MinInt8)),
			}
			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		x, err := Float32(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := int8(x)

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestFloat32FromInt16(t *testing.T) {

	tests := []struct{
		Value int16
	}{
		{
			Value: math.MinInt16,
		},
		{
			Value: -1,
		},
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: math.MaxInt16,
		},
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			test := struct{
				Value int16
			}{
				Value: int16(randomness.Int63n(math.MaxInt16)),
			}
			tests = append(tests, test)

			test = struct{
				Value int16
			}{
				Value: -int16(randomness.Int63n(-1*math.MinInt16)),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		x, err := Float32(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := int16(x)

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestFloat32FromUint8(t *testing.T) {

	tests := []struct{
		Value uint8
	}{
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: math.MaxUint8,
		},
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			test := struct{
				Value uint8
			}{
				Value: uint8(randomness.Int63n(math.MaxUint8)),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		x, err := Float32(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := uint8(x)

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestFloat32FromUint16(t *testing.T) {

	tests := []struct{
		Value uint16
	}{
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: math.MaxUint16,
		},
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			test := struct{
				Value uint16
			}{
				Value: uint16(randomness.Int63n(math.MaxUint16)),
			}
			tests = append(tests, test)
		}
	}

	for testNumber, test := range tests {

		x, err := Float32(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := uint16(x)

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}

}
