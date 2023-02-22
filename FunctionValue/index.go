package main

import "fmt"

func goodBye(name string) string {
	return "Good bye " + name
}

func main(){
	say := goodBye
	result := say("farid")
	 
	fmt.Println(result)
}