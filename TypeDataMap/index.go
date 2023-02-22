package main

import "fmt"

func main(){
	var person map[string]string = map[string]string{
		"firstName" : "farid",
		"lastName" : "fadilah",
	}

	person["work"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])

	book := make(map[string]string)
	book["name"] = "belajar golang"
	book["author"] = "orang google"
	book["ups"] = "salah"

	fmt.Println(book)
	fmt.Println(len(book))
	
	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}