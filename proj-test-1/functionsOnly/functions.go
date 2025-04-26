package functionsonly

import (
	"fmt"
)

func CheckMultiReturnValue() (string, error) {
	some_name := "some_name"
	error1 := fmt.Errorf("norgate math again: square root of negative number")

	return some_name, error1
}


func CheckStrangePrint() {
	age := 35
	name := "shaun"
	s1 := []string{"a", "b", "c"}


	fmt.Printf("%#[3]v\n", age, name, s1 )
}

