package main

import "fmt"

func sayHello(){
	for i := 0; i < 10; i++{
		fmt.Println("hello world")
	}
}

func main(){
	sayHello()
}