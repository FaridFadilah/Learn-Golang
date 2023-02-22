package main

import "fmt"

func main(){
	var(
		counter = 1
		slice = []string{"Farid", "Fadilah", "Permana"}
		person = make(map[string]string)
	)
	
	// FOR LOOP WITHOUT STATEMENT
	for counter <= 5{
		fmt.Println("perulangan ke", counter)
		counter++
	}
	
	// FOR LOOP STATEMENT
	for count := 1; count <= 10; count++{
		fmt.Println("perulangan ke", count)
		} 
	
	// ITERATION ARRAY WITHOUT FOR RANGE
	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	// ITERATION ARRAY WITH FOR RANGE
	for i, value := range slice{
	// for _, value := range slice{
		fmt.Println("index", i, "=", value)
	}

	person["name"] = "farid fadilah permana"
	person["address"] = "cileunyi"
	person["work"] = "programmer"

	for key, value := range person{
		fmt.Println(key, "=", value  )
	}
}