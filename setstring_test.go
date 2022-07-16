package cast_test

import (
	"github.com/reiver/go-cast"
	"github.com/reiver/go-cast/internal/types/stringlike"

	"bytes"
	"database/sql"
	"io"
	"reflect"
	"strings"

	"testing"
)

func TestSetString(t *testing.T) {

	tests := []struct{
		Src string
	}{
		{
			Src: "",
		},



		{
			Src: "apple",
		},
		{
			Src: "banana",
		},
		{
			Src: "cherry",
		},



		{
			Src: "😏 💀👻👽 😊 😍 🙂🙁",
		},



		{
			Src: "😏",
		},
		{
			Src: "💀",
		},
		{
			Src: "👻",
		},
		{
			Src: "👽",
		},
		{
			Src: "😊",
		},
		{
			Src: "😍",
		},
		{
			Src: "🙂",
		},
		{
			Src: "🙁",
		},




		{
			Src: "ONE",
		},
		{
			Src: "TWO",
		},
		{
			Src: "THREE",
		},



		{
			Src: "ONE TWO THREE",
		},



		{
			Src: "۰",
		},
		{
			Src: "۱",
		},
		{
			Src: "۲",
		},
		{
			Src: "۳",
		},
		{
			Src: "۴",
		},
		{
			Src: "۵",
		},
		{
			Src: "۶",
		},
		{
			Src: "۷",
		},
		{
			Src: "۸",
		},
		{
			Src: "۹",
		},



		{
			Src: "۰۱۲۳۴۵۶۷۸۹",
		},
	}

	for testNumber, test := range tests {

		{
			var buffer strings.Builder
			var dst io.Writer = &buffer

			err := cast.SetString(dst, test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %v", err, err)
				continue
			}

			var expected string = test.Src
			var actual   string = buffer.String()

			if expected != actual {
				t.Errorf("For test #%d, actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}

		{
			var actual stringlike.Setter
			var dst interface{SetString(string)} = &actual

			err := cast.SetString(dst, test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %v", err, err)
				continue
			}

			var expected stringlike.Setter = stringlike.SomeSetter(test.Src)

			if expected != actual {
				t.Errorf("For test #%d, actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}

		{
			var actual stringlike.FallibleSetter
			var dst interface{SetString(string)bool} = &actual

			err := cast.SetString(dst, test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %v", err, err)
				continue
			}

			var expected stringlike.FallibleSetter = stringlike.SomeFallibleSetter(test.Src)

			if expected != actual {
				t.Errorf("For test #%d, actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}

		{
			var actual stringlike.ErrableSetter
			var dst interface{SetString(string)error} = &actual

			err := cast.SetString(dst, test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %v", err, err)
				continue
			}

			var expected stringlike.ErrableSetter = stringlike.SomeErrableSetter(test.Src)

			if expected != actual {
				t.Errorf("For test #%d, actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}

		{
			var actual stringlike.Flag
			var dst interface{Set(string)error} = &actual

			err := cast.SetString(dst, test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %v", err, err)
				continue
			}

			var expected stringlike.Flag = stringlike.SomeFlag(test.Src)

			if expected != actual {
				t.Errorf("For test #%d, actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}

		{
			var actual sql.NullString
			var dst sql.Scanner = &actual

			err := cast.SetString(dst, test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %v", err, err)
				continue
			}

			var expected sql.NullString = sql.NullString{
				String: test.Src,
				Valid: true,
			}

			if expected != actual {
				t.Errorf("For test #%d, actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				continue
			}
		}

		{
			var actual string
			var dst *string = &actual

			err := cast.SetString(dst, test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %v", err, err)
				continue
			}

			var expected string = test.Src

			if expected != actual {
				t.Errorf("For test #%d, actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}

		{
			var actual []byte
			var dst *[]byte = &actual

			err := cast.SetString(dst, test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %v", err, err)
				continue
			}

			var expected []byte = []byte(test.Src)

			if !bytes.Equal(expected,actual) {
				t.Errorf("For test #%d, actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}

		{
			var actual []rune
			var dst *[]rune = &actual

			err := cast.SetString(dst, test.Src)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %v", err, err)
				continue
			}

			var expected []rune = []rune(test.Src)

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, actual value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				continue
			}
		}
	}
}

func TestSetString_errorCannot(t *testing.T) {

	var src string = "apple BANANA Cherry 😏 💀👻👽 😊 😍 🙂🙁 ۰۱۲۳۴۵۶۷۸۹"

	var actual int64
	var dst any = &actual

	err := cast.SetString(dst, src)
	if nil == err {
		t.Error("Expected an error, but actually got one.")
		return
	}
	switch err.(type) {
	case cast.CannotSetComplainer:
		// This is the error we expect.
	default:
		t.Errorf("The actual error type is not what was expected.")
		t.Logf("EXPECTED: %s", "cast.CannotCastComplainer")
		t.Logf("ACTUAL:  (%T) %s", err, err)
		return
	}
}

func TestSetString_errorCouldNot(t *testing.T) {

	var src string = "apple BANANA Cherry 😏 💀👻👽 😊 😍 🙂🙁 ۰۱۲۳۴۵۶۷۸۹"

	var actual *stringlike.ErrableSetter = nil
	var dst any = actual

	err := cast.SetString(dst, src)
	if nil == err {
		t.Error("Expected an error, but actually got one.")
		return
	}
	switch err.(type) {
	case cast.CouldNotSetComplainer:
		// This is the error we expect.
	default:
		t.Errorf("The actual error type is not what was expected.")
		t.Logf("EXPECTED: %s", "cast.CouldNotCastComplainer")
		t.Logf("ACTUAL:  (%T) %s", err, err)
		return
	}
}
