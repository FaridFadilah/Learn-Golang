package main

import "fmt"

type Address struct{
	kota, provinsi, negara string
}

func main(){
	address1 := Address{"bandung","jawa barat", "indonesia"}
	var passValue Address = address1 //
	var pointer *Address = &address1 // Pointer
	// pointer := &address1  (ini shortcut mendefinisikan variable pointer)

	passValue.kota = "subang"
	pointer = &Address{"Jakarta", "Jawa barat", "indonesia"}

	fmt.Println(address1) // nilai tidak berubah
	fmt.Println(passValue)
	fmt.Println(pointer)
}