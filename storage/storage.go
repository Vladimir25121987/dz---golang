package storage

import (
	"fmt"
	"os"
)

func ReadBin(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err !=nil {
		return nil, err
	}
	return data, nil
}
func WriteBin(content []byte, name string) {
	file, err := os.Create(name)
	if err !=nil {
		fmt.Println("Ошибка!")
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Println("Ошибка!")
		return
	}
	fmt.Println("Запись успешна!")
	
}