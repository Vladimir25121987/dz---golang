package main

import "time"

type Bin struct {
	id        	string
	private   	bool
	createdAt 	time.Time
	name 		string
}

type BinList struct{
	Bins 		[]Bin
}

func main() {

}

func createBin(private bool, name string, id string) Bin{
	return Bin{
	id:        	id,
	private:   	private,
	createdAt: 	time.Now(),
	name: 		name,
	}
}