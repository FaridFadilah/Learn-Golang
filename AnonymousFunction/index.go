package main

import "fmt"

type BlackList func(string) bool

func registUser(name string, blackList BlackList){
	if blackList(name){
		fmt.Println("You're blocked the name ", name)
	} else{
		fmt.Println("Welcome", name)
	}
}

func main(){
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registUser("root", blacklist)
	registUser("admin", func(name string) bool {
		return name == "admin"
	})
}