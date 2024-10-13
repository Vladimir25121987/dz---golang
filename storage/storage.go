package storage

import (
	"demo/password/bins"
	"demo/password/file"
	"encoding/json"
	"fmt"
	"time"
)

type Storage struct {
	Bins     []bins.Bin `json:"bins"`
	UpdateAt time.Time  `json:"updateAt"`
}

func (storage *Storage) AddBin(bin bins.Bin) {
	storage.Bins = append(storage.Bins, bin)
}

func (storage *Storage) Save() {
	storage.UpdateAt = time.Now()
	data, err := storage.ToBytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteFile(data, "storage.json")
}

func (storage *Storage) ToBytes() ([]byte, error) {
	data, err := json.Marshal(storage)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func NewStorage() (*Storage, error) {
	data, err := file.ReadFile("storage.json")
	if err != nil {
		fmt.Println(err)
		return &Storage{}, nil
	}
	storage := &Storage{}
	err = json.Unmarshal(data, storage)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return storage, nil
}