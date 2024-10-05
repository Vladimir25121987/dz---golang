package bins

Type Bin struct {
	id        	string
	private   	bool
	createdAt 	time.Time
	name 		string
}

Type BinList struct{
	Bins 		[]Bin
}