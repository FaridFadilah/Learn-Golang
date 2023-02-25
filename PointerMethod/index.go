package main

import "fmt"

type Man struct{
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
	// fmt.Println("Di method married : ", man.Name)
}

func main(){
	user := Man{"farid"}
	user.Married()

	fmt.Println(user)
}