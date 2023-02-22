package main

import "fmt"

type Filter func(string) string

func getHelloByFilter(name string, filter Filter){
	nameFilter := filter(name) 
	fmt.Println("Hello",nameFilter)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main(){
	getHelloByFilter("farid", spamFilter)
	getHelloByFilter("anjing", spamFilter)
}