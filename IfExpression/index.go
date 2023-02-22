package main

import "fmt"

func main(){
	name := "farid"

	if name == "f" {
		fmt.Println(name)
	} else if name == "farid"{
		fmt.Println("namaku", name)
	} else{
		fmt.Println("hii")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah sesuai")
	}
}