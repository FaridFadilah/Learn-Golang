package main

import "fmt"

func newMap(name string) map[string]string{
	if name == ""{
		return nil
	} else{
		return map[string]string{
			name: name,
		}
	}
}

func main(){
	var person map[string]string = nil

	if person == nil{
		fmt.Println("kosong")
	}

	newMap("")
}