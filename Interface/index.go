package main

import "fmt"

type HashName interface{
	GetName() string
}

type Person struct{
	Name string
}

type Animal struct{
	Name string
}

func sayHello(hashname HashName){
	fmt.Println("Hello", hashname.GetName())
}

func (animal Animal) GetName() string{
	return animal.Name
}

func (person Person) GetName() string{
	return person.Name
}

func main(){	
	var user Person
	var pet Animal

	user.Name = "farid"
	pet.Name = "cat"

	sayHello(user)
	sayHello(pet)
}