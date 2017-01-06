package cast

import (
	"time"

	"testing"
)

func TestTimeFromTime(t *testing.T) {

	tests := []struct{
		Value time.Time
	}{}
	for n:=0; n<256; n++ {
		test := struct{
			Value time.Time
		}{
			Value: time.Now().Add(time.Duration(n) * -1 * time.Hour),
		}
		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		x, err := Time(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := time.Time(x)

		if expected, actual := test.Value, y; !expected.Equal(actual) {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}
