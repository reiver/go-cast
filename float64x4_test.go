package cast

import (
	"math"

	"testing"
)

func TestFloat64x4FromFloat32x4(t *testing.T) {

	tests := []struct{
		Value [4]float32
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

		const numRand = 5
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
				for _, z := range f32s {
					for _, w := range f32s {

						test := struct{
							Value [4]float32
						}{
							Value: [4]float32{x, y, z, w},
						}

						tests = append(tests, test)
					}
				}
			}
		}
	}


	for testNumber, test := range tests {

		x, err := Float64x4(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := [4]float32{float32(x[0]), float32(x[1]), float32(x[2]), float32(x[3])}

		if expected, actual := test.Value[0], y[0]; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
		if expected, actual := test.Value[1], y[1]; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
		if expected, actual := test.Value[2], y[2]; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
		if expected, actual := test.Value[3], y[3]; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestFloat64x4FromFloat64x4(t *testing.T) {

	tests := []struct{
		Value [4]float64
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

		const numRand = 5
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
				for _, z := range f64s {
					for _, w := range f64s {

						test := struct{
							Value [4]float64
						}{
							Value: [4]float64{x, y, z, w},
						}

						tests = append(tests, test)
					}
				}
			}
		}
	}


	for testNumber, test := range tests {

		x, err := Float64x4(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := [4]float64(x)

		if expected, actual := test.Value[0], y[0]; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
		if expected, actual := test.Value[1], y[1]; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
		if expected, actual := test.Value[2], y[2]; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
		if expected, actual := test.Value[3], y[3]; expected != actual {
			if !(math.IsNaN(float64(expected)) && math.IsNaN(float64(actual))) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
