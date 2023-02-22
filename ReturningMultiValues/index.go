package main

import "fmt"

func getMultiName() (string, string, string) {
	return "farid", "fadilah", "permana"
}

func main(){
	firstName, lastName, _ := getMultiName()
	fmt.Println(firstName, lastName)
}