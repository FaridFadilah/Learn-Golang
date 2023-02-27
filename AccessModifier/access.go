package accessmodifier

import "fmt"

var Apps string = "hello world"
var ver int8 = 1

func sayGood(){
	fmt.Println("this function cant use in another package")
}

func SayGoodbye(){
	fmt.Println("this function can use in another function")
}