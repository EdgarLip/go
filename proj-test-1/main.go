package main

import (
	"fmt"

	mapsfunctions "proj-test-1/mapsfunctions"
	utilsfunctions "proj-test-1/utilsfunctions"
)

func main() {
	fmt.Printf("inside pakcage main-> in main function this is where the go starts !\n")
	utilsfunctions.Print_package_name()
	// utilsfunctions.Print_file_name()
	// utilsfunctions.Print_file_name2()
	// utilsfunctions.Print_utils_test_var() // just testing comments here !!!??
	// utilsfunctions.Check_assimernt_if()
	//------maps ------
	mapsfunctions.Basic_map()

}
