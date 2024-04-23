package utilsfunctions

import (
	"fmt"
	"math/rand"
)

func Print_package_name() {
	fmt.Printf("utils_functions ! \n")
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

func Check_assimernt_if_with_random(x int) {
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}
}

func Check_switch_case_v1(x int) {
	switch {
	case x == 10:
		fmt.Println("will not enter here… use x == 5 for the example  ")
	case x == 5:
		fmt.Println("this is true and this is why it will be printed")
		fallthrough
	case x == 3:
		fmt.Println("will be executed even if not true becasue of fallthrough, will u print me =) ")
	default:
		fmt.Println("default will run if no match case was matched.")
	}
}

func AssertionExample1() {
	var i interface{} = "hello" // assign a string "hello" to an empty interface which is called "i".

	s := i.(string)
	fmt.Println(s) //resault:  hello

	s, ok := i.(string)
	fmt.Println(s, ok) //resault: hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) //resault: 0 false

	//f = i.(float64) // resault: panic    -    panic: interface conversion: interface {} is string, not float64
	//fmt.Println(f)
}
