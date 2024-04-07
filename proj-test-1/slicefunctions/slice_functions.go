package slicefunctions

import "fmt"

func For_loop_on_slice() {
	fmt.Println("------------  for loop over slice basic 1 ------------------")
	xi := []int{42, 43, 44, 45, 46, 47}
	for index, val := range xi {
		fmt.Printf("ranging over a slice index %v , value %v \n", index, val)
	}
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
