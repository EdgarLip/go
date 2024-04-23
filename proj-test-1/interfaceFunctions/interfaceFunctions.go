package interfaceFunctions

import "fmt"

func CheckInterfaceVarible() {
	var val interface{} = 42
	v, ok := val.(int)
	if ok {
		fmt.Println("Double the value:", v*2)
	}
}

func checkTypeWithSwitch(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("val is an int:", v)
	case string:
		fmt.Println("val is a string:", v)
	case bool:
		fmt.Println("val is a boolean:", v)
	case map[string]int:
		fmt.Println("val is a map[string]int:", v)
	case Person:
		fmt.Println("val is a Person struct with name:", v.Name)
	case []int:
		fmt.Println("val is a slice of int:", v)
	default:
		fmt.Println("val is of unknown type")
	}
}

type Person struct {
	Name string
	Age  int
}

func checkTypeWithIf(val interface{}) {
	if v, ok := val.(int); ok {
		fmt.Printf("Type: %T, val is int: %v\n", v, v)
	} else if v, ok := val.(string); ok {
		fmt.Printf("Type: %T, val is string: %v\n", v, v)
	} else if v, ok := val.(bool); ok {
		fmt.Printf("Type: %T, val is bool: %v\n", v, v)
	} else if v, ok := val.(map[string]int); ok {
		fmt.Printf("Type: %T, val is map[string]int: %v\n", v, v)
	} else if v, ok := val.([]int); ok {
		fmt.Printf("Type: %T, val is []int: %v\n", v, v)
	} else if v, ok := val.(Person); ok {
		fmt.Printf("val is a struct Person: %+v\n", v)
		fmt.Printf("Type: %T, val is struct Person: %v\n", v, v)
	} else {
		fmt.Printf("val has an unknown type\n")
	}
}

func CheckTypeWithIfInAction() {
	checkTypeWithIf(42)      // val is an int: 42
	checkTypeWithIf("hello") // val is a string: hello
	checkTypeWithIf(true)    // val is a boolean: true

	m := map[string]int{"foo": 1, "bar": 2} // Map type assertion
	checkTypeWithIf(m)                      // val is a map[string]int: map[bar:2 foo:1]

	p1 := Person{
		Name: "Edgar",
		Age:  20,
	}
	checkTypeWithIf(p1) // val is a struct{ name string } with name: John

	sl := []int{1, 2, 3} // Slice type assertion
	checkTypeWithIf(sl)  // val is a slice of int: [1 2 3]

	var x interface{} = complex(1, 2) // Unknown type assertion
	checkTypeWithIf(x)                // val is of unknown type
}

func CheckTypeWithSwitchInAxtion() {
	checkTypeWithSwitch(42)      // val is an int: 42
	checkTypeWithSwitch("hello") // val is a string: hello
	checkTypeWithSwitch(true)    // val is a boolean: true

	m := map[string]int{"foo": 1, "bar": 2} // Map type assertion
	checkTypeWithSwitch(m)                  // val is a map[string]int: map[bar:2 foo:1]

	p1 := Person{
		Name: "Edgar",
		Age:  20,
	}
	checkTypeWithSwitch(p1) // val is a struct{ name string } with name: John

	sl := []int{1, 2, 3}    // Slice type assertion
	checkTypeWithSwitch(sl) // val is a slice of int: [1 2 3]

	var x interface{} = complex(1, 2) // Unknown type assertion
	checkTypeWithSwitch(x)            // val is of unknown type
}
