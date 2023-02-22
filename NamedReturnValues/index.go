package main

import "fmt"

func getFullName()(firstName string, middleName string, lastName string){
	firstName = "farid"
	middleName = "fadilah"
	lastName = "permana"
	return firstName, middleName, lastName
}

func main(){
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)	
}