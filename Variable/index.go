package main

import "fmt"

func main(){
	var name string
	var friendName string

	var (
		firstName = "farid"
		lastName = "Fadilah"
	)

	fmt.Println("ini multiple var", firstName, lastName)

	name = "Farid Fadilah"
	fmt.Println(name)

	name = "Farid Fadilah Permana"
	fmt.Println(name)

	friendName = "ijal"
	fmt.Println(friendName)
  
	var age int8 = 10
	fmt.Println(age)

	country := "Indonesian"
	fmt.Println(country)

	country = "american"
	fmt.Println(country)
}