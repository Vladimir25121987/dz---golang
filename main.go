package main

import ("time"
)


func main() {
	file, err := WriteBin("")

}

func createBin(private bool, name string, id string) Bin {
	return Bin {
	Id:        	id, 
	Private:   	private,
	CreatedAt: 	time.Now(),
	Name: 		name,
	}
}