package main

import "fmt"

func main(){
	type NoPhone string

	var (
		Phone NoPhone = "08291891239"
	)

	fmt.Println(Phone)
	fmt.Println(NoPhone("08290001923"))
}