package cast_test

import (
	"fmt"

	"github.com/reiver/go-cast"
)

func ExampleBool_convertFromFalseBool() {

	v := false

	b, err := cast.Bool(v)
	if nil != err {
		fmt.Printf("Problem casting to bool: %v \n", err)
		return
	}

	fmt.Printf("b = %t \n", b)

	// Output:
	//
	// b = false
}

func ExampleBool_convertFromTrueBool() {

	v := true

	b, err := cast.Bool(v)
	if nil != err {
		fmt.Printf("Problem casting to bool: %v \n", err)
		return
	}

	fmt.Printf("b = %t \n", b)

	// Output:
	//
	// b = true
}
