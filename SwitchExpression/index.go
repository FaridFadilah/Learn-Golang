package main

import "fmt"

func main(){
	name := "farid"

	switch name{
		case "fadilah":
			fmt.Println("hii fadilah")
		case "permana":
			fmt.Println("hii permana")
		case "farid":
			fmt.Println("hii farid")
		default:
			fmt.Println("Hii, boleh kenalan")
	}

	// SHORT STATEMENT
	switch length := len(name); length > 5{
		case true :
			fmt.Println("benar")
		case false :
			fmt.Println("salah")
	}

	// Switch tanpa kondisi
	length := len(name)
	switch {
		case length > 10:
			fmt.Println("benar")
		case length > 5:
			fmt.Println("salah")
		default:
			fmt.Println("dikatakan salah juga tidak benar")
	}
}