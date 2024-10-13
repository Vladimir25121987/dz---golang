package main

import (
	"demo/password/bins"
	"demo/password/storage"
	"fmt"
)


func main() {
	Bin1, _ := bins.NewBin("1", true, "Bin1")

	Storage1, err := storage.NewStorage()
	if err != nil {
		fmt.Print(err)
	}
	Storage1.AddBin(*Bin1)
	Storage1.Save()
}




