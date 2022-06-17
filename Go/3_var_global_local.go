package main

import "fmt"

var globalVar string = "I'm global var!"

func main() {
	localVar := "I'm local var"
	fmt.Println(globalVar)
	fmt.Println(localVar)
}
