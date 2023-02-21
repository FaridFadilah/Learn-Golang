package main

import "fmt"

func main(){
	var (
		month = [...]string{
			"Januari",
			"Februari",
			"Maret",
			"April",
			"Mei",
			"Juni",
			"Juli",
			"Agustus",
			"September",
			"Oktober",
			"November",
			"Desember",
		}

		slice1 = month[4:7]
	)

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	
	month[5] = "dirubah"
	fmt.Println(slice1)

	slice1[0] = "dirubah"
	fmt.Println(slice1)
}