package main

import "fmt"

func sumAll(numbers ...int) int {
	result := 0
	for _, value := range numbers{
		result += value
	}
	return result
}

func main(){
	total := sumAll(10,10,10,10)
	fmt.Println(total)

	number := []int{10,20,30,40,50}
	total = sumAll(number...)
	fmt.Println(total)
}