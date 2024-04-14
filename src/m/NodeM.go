package m

import (
	Data "Golang/src/data"
	"Golang/src/node2"
	"fmt"
)

var ch = make(chan Data.Data)

type NodeM interface {
	Run(node1 chan Data.Data)
	Data.Channel
}

func (data data.Data) Run(node1Ch chan data.Data) {
	node1Ch <- data
	node2 = <-ch
	append(node2.Node2, "")
	fmt.Println(node2.Node2)

	// send back data to Node1
	ch1 := make(chan data.Data)
	go node2.Run(ch1)
}

func GetChannel() chan Data.Data {
	return ch
}
