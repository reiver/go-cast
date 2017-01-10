package cast_test

import (
	"fmt"

	"github.com/reiver/go-cast"
)

type FuzzyLogic struct {
	value float64
}

func (receiver FuzzyLogic) Bool() (bool, error) {
	result := receiver.value > 0.5

	return result, nil
}

var (
	FuzzyLogicFalse = FuzzyLogic{value: 0.0}
	FuzzyLogicMaybe = FuzzyLogic{value: 0.5}
	FuzzyLogicTrue  = FuzzyLogic{value: 1.0}
)

func ExampleBool_customTypeFuzzyLogic1() {

	v := FuzzyLogicFalse

	b, err := cast.Bool(v)
	if nil != err {
		fmt.Printf("Problem casting to bool: %v \n", err)
		return
	}

	fmt.Printf("b = %t\n", b)

	// Output:
	//
	// b = false
}

func ExampleBool_customTypeFuzzyLogic2() {

	v := FuzzyLogicTrue

	b, err := cast.Bool(v)
	if nil != err {
		fmt.Printf("Problem casting to bool: %v \n", err)
		return
	}

	fmt.Printf("b = %t\n", b)

	// Output:
	//
	// b = true
}
