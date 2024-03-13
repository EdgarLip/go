package main

import (
	"fmt"

	mapsfunctions "edgar.com/proj-test-1/mapsfunctions"
	utilsfunctions "edgar.com/proj-test-1/utilsfunctions"

	//pt2helperfunctions "edgar.com/proj-test-2/pt2helperfunctions"
	mygoutil "github.com/EdgarLip/mygoutils"
)

func main() {
	fmt.Printf("inside pakcage main-> in main function this is where the go starts !\n")
	//--- import package in same module ---
	utilsfunctions.Print_package_name()
	// utilsfunctions.Print_file_name()
	// utilsfunctions.Print_file_name2()
	// utilsfunctions.Print_utils_test_var() // just testing comments here !!!??
	// utilsfunctions.Check_assimernt_if()
	//------maps ------
	mapsfunctions.Basic_map()

	//--- import module ---
	//fmt.Printf("other package name = %s", pt2helperfunctions.Get_package_name())
	fmt.Println(mygoutil.Get_package_name())

}
