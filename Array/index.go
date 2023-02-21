package main

import "fmt"

func main(){
	var names [3]string
		names[0] = "farid"
		names[1] = "fadilah"
		names[2] = "permana"
		
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	
	var arr = [3]int{1,2,3,}

	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	fmt.Println(len(names))
	fmt.Println(len(arr))
}