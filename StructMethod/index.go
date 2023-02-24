package main

import "fmt"

type Customer struct{
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string){
	fmt.Println("hello", name, "my name", customer.Name)
}

func (a Customer) sayHii(name string){
	fmt.Println("Hiii", name)
}

func main(){
	var user Customer

	user.Name = "Farid" 
	user.Address = "Cileunyi" 
	user.Age = 19

	user.sayHello("fadilah")
	user.sayHii("permana")

}