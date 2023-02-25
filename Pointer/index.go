package main

import "fmt"

type Address struct{
	kota, provinsi, negara string
}

func main(){
	value := Address{"bandung","jawa barat", "indonesia"}
	// var passValue Address = value 
	var pointer *Address = &value // Pointer
	
	// pointer := &value  (ini shortcut mendefinisikan variable pointer)
	pointerNew := new(Address) // membuat pointer menggunakan function new

	pointerNew.kota = "sumedang"
	pointerNew.provinsi = "jawa barat"
	pointerNew.negara = "indonesia"

	// passValue.kota = "subang"
	pointer.kota = "jogja"

	fmt.Println(value)
	fmt.Println(pointer)

	// pointer = &Address{"jepara", "jawa tengah", "indonesia"}

	// fmt.Println(value) // tidak berubah
	// fmt.Println(passValue)
	// fmt.Println(pointer)

	*pointer = Address{"malang", "jawa timur", "indonesia"}

	fmt.Println(value)
	fmt.Println(pointer)
	fmt.Println(pointerNew)
}