package functionsonly

import (
	"fmt"
)

func CheckMultiReturnValue() (string, error) {
	some_name := "some_name"
	error1 := fmt.Errorf("norgate math again: square root of negative number")

	return some_name, error1
}



