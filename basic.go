package main

import "fmt"

//func variableZeroValue() {
//	var a int
//	var s string
//}

var (
	test1 = 1
	test2 = "2"
)

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c = 3, 4, false
	var s = "abc"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "abc"
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println(" Hello world ")
	//variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(test1, test2)
}
