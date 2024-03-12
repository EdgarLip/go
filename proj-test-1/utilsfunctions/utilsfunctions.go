package utilsfunctions

import "fmt"

func Print_package_name() {
	fmt.Printf("utils_functions \n")
}

func Print_file_name() {
	fmt.Printf("utils_functions\n")
}

func Print_utils_test_var() {
	fmt.Printf("utils_test_var = %d\n", Utils_helper_var)
}

func Check_assimernt_if() {
	if x := 10; x > 12 {
		fmt.Printf("x = %d, we are in the if statment !!! \n", x)
	} else {
		fmt.Printf("x = %d, we are in else  statment !!! \n", x)
	}
}
