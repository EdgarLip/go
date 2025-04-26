package slicefunctions

import (
	"fmt"
	"strings"
)

// For_loop_on_slice Function is nice and good example of how the basic for loop works
// in the go programing language.
func For_loop_on_slice() {
	fmt.Println("------------  for loop over slice basic 1 ------------------")
	xi := []int{42, 43, 44, 45, 46, 47}
	for index, val := range xi {
		fmt.Printf("ranging over a slice index %v , value %v \n", index, val)
	}
}


func JoinStringsInSlice(_strings []string, seperator string) string {
	return strings.Join(_strings, seperator)
}


func TestCopy() {
	// in this example is am showing that nested copy does not work at all for a slice.
	fmt.Println("---- Running TestCopy ---- ")
	matrix := [][]int{{1, 2, 3}, {1, 2, 3}, {1}} //  with values
	matrix_cpy := make([][]int, 0, 10)
	copy(matrix_cpy, matrix)

	matrix[0][0] = 202

	fmt.Println(matrix)     //[[1 2 3] [1 2 3] [1]]
	fmt.Println(matrix_cpy) //[]
	fmt.Println("---- End TestCopy ---- ")
}

func CheckVeriadic(x ...int) { //  function is getting an unlimited number of ints
	fmt.Println(" --- CheckVeriadic --- ")
	fmt.Println(x)      // resault: [1 2 3 4 5 6 7]
	fmt.Printf("%T \n", x) //resault:  []int                 , this is the type of the ...int
	fmt.Println(" --- CheckVeriadic end --- ")
}


func SliceTheSlice() {
	fmt.Println(" --- SliceTheSlice --- ")
	s1 := []int{1, 2, 3, 4, 5}
    s2 := s1[0:2:2]
	fmt.Printf("s2 = %v \n", s2)                 // s2 = [1,2]
	fmt.Printf("s1 pointer = %p \n", &s1)        // res: s1 pointer = 0xc000008210
	fmt.Printf("s2 pointer = %p \n", &s2)        // res: s2 pointer = 0xc000008228

	fmt.Printf("s1 cap = %v \n", cap(s1))
	fmt.Printf("s1 len = %v \n", len(s1))
	fmt.Printf("s2 cap = %v \n", cap(s2))
	fmt.Printf("s2 len = %v \n", len(s2))

	fmt.Printf("s1[0] pointer = %p \n", &s1[0])  // res: s1[0] pointer = 0xc0000108a0
	fmt.Printf("s2[0] pointer = %p \n", &s2[0])  // res: s2[0] pointer = 0xc0000108a0

	s2[0] = 200
	fmt.Printf("s1 = %v \n", s1)      // s1 = [200 2 3 4 5]
	fmt.Printf("s2 = %v \n", s2)      // s2 = [200 2]
	s2 = append(s2, 500)
	fmt.Printf("s1 = %v \n", s1)      // s1 = [200 2 3 4 5
	fmt.Printf("s2 = %v \n", s2)      // s2 = [200 2 500]

	fmt.Printf("s1[0] pointer = %p \n", &s1[0])   // s1[0] pointer = 0xc0000108a0
	fmt.Printf("s2[0] pointer = %p \n", &s2[0])   // s2[0] pointer = 0xc000016320
}