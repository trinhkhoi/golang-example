package data

type Data struct {
	Id    string
	Node1 bool
	Node2 []string
	Node3 bool
}

type Channel interface {
	GetChannel() chan Data
}
