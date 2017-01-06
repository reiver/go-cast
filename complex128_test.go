package cast

import (
	"math"

	"testing"
)

func TestComplex128FromComplex64(t *testing.T) {

	tests := []struct{
		Value complex64
	}{}
	{
		f32s := []float32{
			float32(math.Inf(-1)),
			-math.MaxFloat32,
			-math.Pi,
			-math.E,
			-math.Sqrt2,
			-1.0,
			-math.Ln2,
			-math.SmallestNonzeroFloat32,
			0.0,
			math.SmallestNonzeroFloat32,
			math.Ln2,
			1.0,
			math.Sqrt2,
			math.E,
			math.Pi,
			math.MaxFloat32,
			float32(math.Inf(+1)),

			float32(math.NaN()),
		}

		const numRand = 20
		for i:=0; i<numRand; i++ {
			f32 := randomness.Float32()
			f32s = append(f32s, f32)

			f32  = -randomness.Float32()
			f32s = append(f32s, f32)



			f32  = randomness.Float32() * math.MaxFloat32
			f32s = append(f32s, f32)

			f32  = -randomness.Float32() * math.MaxFloat32
			f32s = append(f32s, f32)



			f32  = randomness.Float32() * 999999999
			f32s = append(f32s, f32)

			f32  = -randomness.Float32() * 999999999
			f32s = append(f32s, f32)
		}

		for _, x := range f32s {
			for _, y := range f32s {

				test := struct{
					Value complex64
				}{
					Value: complex(x, y),
				}

				tests = append(tests, test)
			}
		}
	}


	for testNumber, test := range tests {

		x, err := Complex128(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := complex64(x)

		if expected, actual := real(test.Value), real(y); expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
		if expected, actual := imag(test.Value), imag(y); expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestComplex128FromComplex128(t *testing.T) {

	tests := []struct{
		Value complex128
	}{}
	{
		f64s := []float64{
			math.Inf(-1),
			-math.MaxFloat64,
			-math.Pi,
			-math.E,
			-math.Sqrt2,
			-1.0,
			-math.Ln2,
			-math.SmallestNonzeroFloat64,
			0.0,
			math.SmallestNonzeroFloat64,
			math.Ln2,
			1.0,
			math.Sqrt2,
			math.E,
			math.Pi,
			math.MaxFloat64,
			math.Inf(+1),

			math.NaN(),
		}

		const numRand = 20
		for i:=0; i<numRand; i++ {
			f64 := randomness.Float64()
			f64s = append(f64s, f64)

			f64  = -randomness.Float64()
			f64s = append(f64s, f64)



			f64  = randomness.Float64() * math.MaxFloat64
			f64s = append(f64s, f64)

			f64  = -randomness.Float64() * math.MaxFloat64
			f64s = append(f64s, f64)



			f64  = randomness.Float64() * 999999999
			f64s = append(f64s, f64)

			f64  = -randomness.Float64() * 999999999
			f64s = append(f64s, f64)
		}

		for _, x := range f64s {
			for _, y := range f64s {

				test := struct{
					Value complex128
				}{
					Value: complex(x, y),
				}

				tests = append(tests, test)
			}
		}
	}


	for testNumber, test := range tests {

		x, err := Complex128(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := complex128(x)

		if expected, actual := real(test.Value), real(y); expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
		if expected, actual := imag(test.Value), imag(y); expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestComplex128FromFloat32(t *testing.T) {

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

		x, err := Complex128(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		yi := imag(x)
		if expected, actual := float64(0), yi; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}

		y := float32(real(x))

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestComplex128FromFloat64(t *testing.T) {

	tests := []struct{
		Value float64
	}{
		{
			Value: math.Inf(-1),
		},
		{
			Value: -math.MaxFloat64,
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
			Value: -math.SmallestNonzeroFloat64,
		},
		{
			Value: 0.0,
		},
		{
			Value: math.SmallestNonzeroFloat64,
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
			Value: math.MaxFloat64,
		},
		{
			Value: math.Inf(+1),
		},



		{
			Value: math.NaN(),
		},
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			test := struct{
				Value float64
			}{
				Value: randomness.Float64(),
			}
			tests = append(tests, test)

			test = struct{
				Value float64
			}{
				Value: -randomness.Float64(),
			}
			tests = append(tests, test)



			test = struct{
				Value float64
			}{
				Value: randomness.Float64() * math.MaxFloat64,
			}
			tests = append(tests, test)

			test = struct{
				Value float64
			}{
				Value: -randomness.Float64() * math.MaxFloat64,
			}
			tests = append(tests, test)



			test = struct{
				Value float64
			}{
				Value: randomness.Float64() * 999999999,
			}
			tests = append(tests, test)

			test = struct{
				Value float64
			}{
				Value: -randomness.Float64() * 999999999,
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		x, err := Complex128(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		yi := imag(x)
		if expected, actual := float64(0), yi; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}

		y := float64(real(x))

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestComplex128FromInt8(t *testing.T) {

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

		x, err := Complex128(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		yi := imag(x)
		if expected, actual := float64(0), yi; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}

		y := int8(real(x))

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestComplex128FromInt16(t *testing.T) {

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

		x, err := Complex128(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		yi := imag(x)
		if expected, actual := float64(0), yi; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}

		y := int16(real(x))

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestComplex128FromUint8(t *testing.T) {

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

		x, err := Complex128(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		yi := imag(x)
		if expected, actual := float64(0), yi; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}

		y := uint8(real(x))

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestComplex128FromUint16(t *testing.T) {

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

		x, err := Complex128(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		yi := imag(x)
		if expected, actual := float64(0), yi; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}

		y := uint16(real(x))

		if expected, actual := test.Value, y; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
