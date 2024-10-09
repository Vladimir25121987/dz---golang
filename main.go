package main

import (
	"demo/password/bins"
	"demo/password/storage"
	"fmt"
	"time"
)


func main() {
	
	file, err := storage.NewStorage()
	if err !=nil {
		fmt.Println("Ошибка!")
	}
	return file, nil



}

func CreateBin(private bool, name string, id string) bins.Bin {
	return bins.Bin {
	Id:        	id, 
	Private:   	private,
	CreatedAt: 	time.Now(),
	Name: 		name,
	}
}