package stringfunctions


import (
	"fmt"
	"strings"
)


func CheckSplitFunc() {
	str1 := "some text around here"
	fmt.Println(strings.Split(str1, " "))
	fmt.Println(strings.Split(str1, "some"))
	fmt.Println(strings.Split(str1, "text"))
}