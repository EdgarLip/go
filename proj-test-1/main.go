package main

import (
	"fmt"

	utilsfunctions "proj-test-1/utilsfunctions"
)

func main() {
	fmt.Printf("inside pakcage main-> in main function this is where the go starts !\n")
	utilsfunctions.Print_package_name()
	utilsfunctions.Print_file_name()
	utilsfunctions.Print_file_name2()
	utilsfunctions.Print_utils_test_var()
}

//	utils "proj-test-1/utils_functions/utils_functions.go"
