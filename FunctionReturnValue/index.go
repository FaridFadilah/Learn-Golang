package main

import "fmt"

func getHello(name string) string {
	if len(name) <= 1 {
		return "nama tidak boleh kosong"
	} else{
		return "hallo " + name
	}
}

func main(){
	resultNull := getHello("")
	result := getHello("farid")
	fmt.Println(resultNull)
	fmt.Println(result)
}