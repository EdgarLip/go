package closurefunctions

import (
	"fmt"
)

func executeFuncD(d func()) {
	d()
}

// funtion resault:
// ---------------
// 0 0xc00000a140
// 1 0xc00000a140
// 2 0xc00000a140
// 3 0xc00000a140
func DoTheThing1() {
	fmt.Println("--- Running DoTheThing1() --- ")
	for i := 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d %p\n", i, &i)
		}
		executeFuncD(v)
	}
}

// funtion resault:
// ---------------
// 4 0xc00000a148
// 4 0xc00000a148
// 4 0xc00000a148
// 4 0xc00000a148
func DoTheThing2() {
	fmt.Println("--- Running DoTheThing2() --- ")
	s := make([]func(), 4)
	for i := 0; i < 4; i++ {
		s[i] = func() {
			// they all point to the same "i"
			fmt.Printf("%d %p\n", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}
}

// funtion resault:
// ---------------
// 0 0xc00000a150
// 1 0xc00000a158
// 2 0xc00000a160
// 3 0xc00000a168
func DoTheThing2WithSolution() {
	fmt.Println("--- Running DoTheThing2WithSolution --- ")
	s := make([]func(), 4)
	for i := 0; i < 4; i++ {
		j := i
		s[i] = func() {
			// they all point to the same "i"
			fmt.Printf("%d %p\n", j, &j)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}
}
