package main

import "fmt"

func apaAja(i int) interface{}{
	if i == 1{
		return 1
	} else if i == 2{
		return true
	} else{
		return "ups"
	}
}

func main(){
	null := apaAja(1)

	fmt.Println(null)
}