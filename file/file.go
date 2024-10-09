package file

import (
	"fmt"
	"os"
)


func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err !=nil {
		return nil, err
	}
	return data, nil
}
func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err !=nil {
		fmt.Println("Ошибка!")
	}
	_, err = file.Write(content)
	defer file.Close()
	if err !=nil {
		fmt.Println("Ошибка!")
		return
	}
	fmt.Println("Запись успешна!")
}