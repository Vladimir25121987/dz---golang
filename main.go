package main

import "time"

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