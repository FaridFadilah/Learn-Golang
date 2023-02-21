package main

import "fmt"

func main(){
	var (
		firstName = "farid"
		middleName = "farid"
		lastName = "fadilah"
		resultTrue = firstName == middleName
		resultFalse = firstName == lastName
		value = 100
		number = 200
	)

	fmt.Println(resultTrue)
	fmt.Println(resultFalse)
	fmt.Println(value > number)
	fmt.Println(value < number)
	fmt.Println(value >= number)
	fmt.Println(value <= number)
	fmt.Println(value == number)
	fmt.Println(value != number)
}