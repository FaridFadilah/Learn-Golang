package main

import "fmt"

type Address struct{
	kota, provinsi, negara string
}

// func changeCountryToIndonesia(address Address){  // ini bukan function pointer
func changeCountryToIndonesia(address *Address){ // ini function pointer
	address.negara = "indonesia"
}

func main(){
	var pointer Address = Address{
		kota: "bandung",
		provinsi: "jawa barat",
		negara: "malaysia",
	}

	var pointerParam *Address = &pointer
	changeCountryToIndonesia(pointerParam)
	fmt.Println(pointer)
}