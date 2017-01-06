package cast_test

import (
	"fmt"

	"github.com/reiver/go-cast"
)

// Modular3 is an implementation of "mod 3" arithmetic system.
//
// Modular arithmetic systems "wrap around".
//
// You can think of them working like a clock or an odometer digit.
//
// So, with "mod 3" the sequence of numbers is:
// 0, 1, 2, 0, 1, 2, 0, 1, 2, ....
//
// Or, in other words:
//
//	0 + 1 = 1
//
//	1 + 1 = 2
//
//	2 + 1 = 0
//
// For the Modular3 type, adding 1 (i.e., "+ 1") capabilities is provided
// by the Inc method.
type Modular3 struct {
	value uint8
}

// Inc adds 1 (i.e., "+ 1") to a Modular3.
func (receiver *Modular3) Inc(){
	receiver.value = (receiver.value + 1) % 3
}

// String makes Modular3 fit the fmt.Stinger interface,
// so that it will have "nice" output when using it
// with (for example) funcs such as fmt.Printf().
func (receiver Modular3) String() string {
	return fmt.Sprintf("%d₃", receiver.value)
}

// Int64 makes Modular3 work with cast.Int64()
func (receiver Modular3) Int64() (int64, error) {
	u8 := receiver.value % 3

	i64 := int64(u8)

	return i64, nil
}

func ExampleInt64_customType() {

	var x Modular3

	fmt.Printf("x = %v\n", x)
	i64, err := cast.Int64(x)
	if nil != err {
		fmt.Printf("Could not convert to int64, because: %v", err)
		return
	}
	fmt.Printf("i64 = %v\n\n", i64)

	x.Inc()

	fmt.Printf("x = %v\n", x)
	i64, err = cast.Int64(x)
	if nil != err {
		fmt.Printf("Could not convert to int64, because: %v", err)
		return
	}
	fmt.Printf("i64 = %v\n\n", i64)

	x.Inc()

	fmt.Printf("x = %v\n", x)
	i64, err = cast.Int64(x)
	if nil != err {
		fmt.Printf("Could not convert to int64, because: %v", err)
		return
	}
	fmt.Printf("i64 = %v\n\n", i64)

	x.Inc()

	fmt.Printf("x = %v\n", x)
	i64, err = cast.Int64(x)
	if nil != err {
		fmt.Printf("Could not convert to int64, because: %v", err)
		return
	}
	fmt.Printf("i64 = %v\n\n", i64)

	x.Inc()

	fmt.Printf("x = %v\n", x)
	i64, err = cast.Int64(x)
	if nil != err {
		fmt.Printf("Could not convert to int64, because: %v", err)
		return
	}
	fmt.Printf("i64 = %v\n\n", i64)

	// Output:
	// x = 0₃
	// i64 = 0
	//
	// x = 1₃
	// i64 = 1
	//
	// x = 2₃
	// i64 = 2
	//
	// x = 0₃
	// i64 = 0
	//
	// x = 1₃
	// i64 = 1
}
