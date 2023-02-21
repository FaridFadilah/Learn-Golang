package main

import "fmt"

func main(){
	var (
		x = 2
		y = 3
		i = 10
	)

	// ini operasi matematika biasa
	fmt.Println("penjumlahan dari ", x, y,x+y)
	fmt.Println("pengurangan dari ", x, y,x-y)
	fmt.Println("perkalian dari", x, y, x*y)
	fmt.Println("pembagian dari", y, x, y/x)
	fmt.Println("sisa pembagian dari", y, x, y/x)

	// ini operasi matematika dengan Augmented Assigned
	i += 10
	fmt.Println("ini Augmented Assigned dari penjumlahan", i)
	
	i -= 5
	fmt.Println("ini Augmented Assigned dari pengurangan", i)
	
	i *= 10
	fmt.Println("ini Augmented Assigned dari perkalian", i)

	i /= 2
	fmt.Println("ini Augmented Assigned dari pembagian", i)
	
	i %= 9
	fmt.Println("ini Augmented Assigned dari sisa pembagian", i)
	
	// ini operasi matematika dengan Unary Operation
	i++ 
	fmt.Println("ini penjumlahan dengan Unary Operation", i)
	
	i--
	fmt.Println("ini pengurangan dengan Unary Operation", i)

	
}