package main

import (
	"fmt"

	cuncurrencyFunctions "edgar.com/proj-test-1/cuncurrencyFunctions"
	interfaceFunctions "edgar.com/proj-test-1/interfaceFunctions"
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

// run this file from wsl.
// go to the folder -> /mnt/c/Users/lipnitsk/go_projects/go/proj-test-1
// and run          -> go run main.go
// git work flow basic: 
//   - git checkout -b <branch name > 
//   - git add <files that u need > 
//   - git commit -m "msg"
//   - git push origin <local_baranch_name>:<remote_branch_name>      ->     check that that your branch is uploaded to the reamote repo.
func main() {
	fmt.Println("run this in WSL from the main test with WSL !!!! ")
	fmt.Printf("inside pakcage main-> in main function this is where the go starts !\n")
	//--- import package in same module ---
	//utilsfunctions.Print_package_name()
	utilsfunctions.Print_file_name()
	// utilsfunctions.Print_file_name2()
	// utilsfunctions.Print_utils_test_var() // just testing comments here !!!??
	// utilsfunctions.Check_assimernt_if()

	// --- time ---
	fmt.Println("--- time start ---")
	fmt.Println(utilsfunctions.TimeMesureFunc(utilsfunctions.DoWork))
	fmt.Println("--- time end ---")

	//  --- random ---
	//utilsfunctions.Check_assimernt_if_with_random(10)

	// ------ switch ------
	//utilsfunctions.Check_switch_case_v1(5)

	// ---------------------
	// ------slices ---------
	// ---------------------
	sliceForTest := []int{1, 2, 3, 4, 5, 6}
	slicefunctions.For_loop_on_slice()
	slicefunctions.TestCopy()
	slicefunctions.CheckVeriadic(sliceForTest...)

	// -------------------
	// ------ maps -------
	// -------------------
	fmt.Println("--- start map ---")
	//mapsfunctions.Basic_map()
	//mapsfunctions.For_loop_on_map()
	//mapsfunctions.Basic_unmarshal1()
	//mapsfunctions.Basic_unmarshal2()
	mapsfunctions.InterfacedMapExampleInAction()
	mapsfunctions.InterfacedMapExampleInActionV2()
	mapsfunctions.InterfacedMapExampleInActionV3()
	fmt.Println("--- end map ---")


	
    // -------------------------------------------
	// ------structs | struct section |  ---------
	// -------------------------------------------
	fmt.Println("--- start structs ---")

	bill1 := structfunctions.Bill{Name: "nomis-bday",
	                              Items: map[string]float64{
								  	"slalt": 10.2,
								  	"waffel": 5.3,
								  },
								  Tip: 10.10}

	bill2 := &structfunctions.Bill{Name: "yulias-bday",
	                              Items: map[string]float64{
								  	"slalt": 12.2,
								  	"waffel": 3.3,
								  },
								  Tip: 10.10}

	bill1.UpdateTip(5.5)
	bill2.UpdateTip(5.5)
	bill1.UpdateTipWithPointer(2.5)
	bill2.UpdateTipWithPointer(1.5)
	fmt.Println(bill1.GetBillName())	
	fmt.Println(bill2.GetBillName())	
	//interfaceFunctions.PrintTipsOfanyBillType(bill1)
	interfaceFunctions.PrintTipsOfanyBillType(bill2)


	//structfunctions.ModifyPersonInAction()
	
	fmt.Println("--- end structs ---")

	//---- import module ------
	//fmt.Printf("other package name = %s", pt2helperfunctions.Get_package_name())
	// fmt.Println(mygoutil.Get_package_name())
	mygoutil.Print_tag_v1_2_0()
	structfunctions.PersonSayHiInAction()
	structfunctions.InfoInAction()

	//   ----------------------------------
	//   ------ interface section ---------
	//   ----------------------------------
	fmt.Println("")
	fmt.Println("--- interface start  ---")
	fmt.Println("***\t\t check type with if \t\t***")
	interfaceFunctions.CheckTypeWithIfInAction()
	fmt.Println("***\t\t check type with switch \t\t***")
	interfaceFunctions.CheckTypeWithSwitchInAxtion()

	utilsfunctions.AssertionExample1()
	fmt.Println("--- interface end  ---")
	fmt.Println("")

	//   ------ file section ---------
	// fmt.Println(" --- file work --- ")
	// content := "Hello, world!\nThis is a test file\nfrom Edgar !."

	// opend_file, err := utilsfunctions.CreateFile("text-test.txt")
	// if err != nil {
	//     panic(err)
	// }
	// defer opend_file.Close()

	// //opend_file2, err2 := utilsfunctions.OpenFile("text-test.txt")
	// if err != nil {
	// 	panic(err) // Handle the error if there was a problem opening the file
	// }

	// if err := utilsfunctions.WriteToFile(opend_file, content); err != nil {
	// 	panic(err)
	// }

	// if err := utilsfunctions.FileOffsetToStartOfFile(opend_file); err != nil {
	// 	panic(err)
	// }

	// if err := utilsfunctions.ReadFile(opend_file); err != nil {
	// 	panic(err) // Handle the error if there was a problem reading from the file
	// }
	// fmt.Println(" --- end file work --- ")

	//   -------------------------------
	//   ---   cuncurrency section   ---
	//   -------------------------------
	// fmt.Println(" --- cuncurrency Start --- ")

	cuncurrencyFunctions.CuncurencyPlaceHolderfucntion()
	// cuncurrencyFunctions.CuncurencyInAction()
	// cuncurrencyFunctions.ChannelTestInAction()
	// cuncurrencyFunctions.FanInMain()
	// fmt.Println(" --- cuncurrency End --- ")

}
