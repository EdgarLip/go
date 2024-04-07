package main

import (
	"fmt"

	mapsfunctions "edgar.com/proj-test-1/mapsfunctions"
	slicefunctions "edgar.com/proj-test-1/slicefunctions"
	structfunctions "edgar.com/proj-test-1/structfunctions"
	utilsfunctions "edgar.com/proj-test-1/utilsfunctions"

	//pt2helperfunctions "edgar.com/proj-test-2/pt2helperfunctions"
	mygoutil "github.com/EdgarLip/mygoutils"
)

func init() {
	fmt.Println("INIT function runs first ! ")
}

func main() {
	fmt.Printf("inside pakcage main-> in main function this is where the go starts !\n")
	//--- import package in same module ---
	//utilsfunctions.Print_package_name()
	utilsfunctions.Print_file_name()
	// utilsfunctions.Print_file_name2()
	// utilsfunctions.Print_utils_test_var() // just testing comments here !!!??
	// utilsfunctions.Check_assimernt_if()

	//  --- random ---
	//utilsfunctions.Check_assimernt_if_with_random(10)

	// ------ switch ------
	//utilsfunctions.Check_switch_case_v1(5)

	//------slices ---------
	slicefunctions.For_loop_on_slice()
	slicefunctions.TestCopy()

	//------maps ---------
	mapsfunctions.Basic_map()
	//mapsfunctions.For_loop_on_map()

	// ------structs ---------
	structfunctions.ModifyPersonInAction()
	//---- import module ------
	//fmt.Printf("other package name = %s", pt2helperfunctions.Get_package_name())
	// fmt.Println(mygoutil.Get_package_name())
	mygoutil.Print_tag_v1_2_0()

}
