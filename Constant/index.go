package main

import "fmt"

func main(){
	const name string = "farid fadilah"
	
	const (
		firstName = "Farid"
		lastName = "Fadilah"
	)

	fmt.Println("ini single var const", name)
	fmt.Println("ini multiple const", firstName, lastName)
}