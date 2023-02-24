package main

import "fmt"

func random() interface{} {
	return "10"
}

func main(){
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type){
		case string:
			fmt.Println("ini string", value)
		case int:
			fmt.Println("ini Int", value)
		default:
			fmt.Println("unknown")
	}
	
	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)
}