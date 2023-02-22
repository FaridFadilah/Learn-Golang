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

		day = [6]string{
			"senin",
			"selasa",
			"rabu",
			"kamis",
			"jumat",
			"sabtu",
		}

		iniArray = [...]int{1,2,3,4,5}
		iniSlice = []int{1,2,3,4,5}
		
		sliceDay = day[0:1]
		sliceMonth = month[4:7]
	)
	
	newSlice := make([]string, 2, 5)
	newSlice[0] = "farid"
	newSlice[1] = "fadilah"
	
	fmt.Println("ini membuat slice baru",newSlice)
	fmt.Println("panjang dari slice baru",len(newSlice))
	fmt.Println("capacity dari slice baru",cap(newSlice))
	
	copySlice := make([]string, 1, cap(newSlice))
	
	copy(copySlice, newSlice)

	fmt.Println("penggabungan antara copySlice dengan newSlice", copySlice)

	fmt.Println(sliceMonth)
	fmt.Println(len(sliceMonth))
	fmt.Println(cap(sliceMonth))
	
	month[5] = "dirubah"
	fmt.Println(sliceMonth)

	sliceMonth[0] = "dirubah"
	fmt.Println(sliceMonth)

	tambahHari := append(sliceDay, "minggu")
	fmt.Println(tambahHari)

	tambahHariLagi := append(tambahHari,"senin lagi")
	fmt.Println(tambahHariLagi)

	fmt.Println("ini array", iniArray)
	fmt.Println("ini slice", iniSlice)
}