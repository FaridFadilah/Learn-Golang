package main

import "fmt"

type Customer struct{
	Name, Address string
	Age int
}

func main(){
	var user Customer

	user.Name = "Farid" 
	user.Address = "Cileunyi" 
	user.Age = 19

	fmt.Println(user)
	fmt.Println(user.Name) 
	fmt.Println(user.Address) 
	fmt.Println(user.Age)

	users := Customer{
		Name : "LOL",
		Address: "California",
		Age : 12,
	}

	fmt.Println(users) 
	fmt.Println(users.Name) 
	fmt.Println(users.Address) 
	fmt.Println(users.Age)

	another := Customer{"LOSS", "Jakarta", 10}

	fmt.Println(another)
	fmt.Println(another.Name)
	fmt.Println(another.Address)
	fmt.Println(another.Age)
}