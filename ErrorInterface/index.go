package main

import (
	"errors"
	"fmt"
)

type error interface{
	Error() string
}

func pembagi(nilai int, pembagi int) (int, error){
	if pembagi == 0{
		return 0, errors.New("tidak boleh kosong")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main(){
	var contohError error = errors.New("Error")
	fmt.Println(contohError.Error())

	hasil, err := pembagi(10,2)

	if err == nil{
		fmt.Println("hasil", hasil)
	} else{
		fmt.Println("error", err.Error())
	}
}

