package node3

import (
	"Golang/src/data"
	"fmt"
)

type Node3 interface {
	Run(node1 chan data.Data)
	data.Channel
}

func (data data.Data) Run(node1Ch chan data.Data) {
	node1Ch <- data
}

func (node2 data.Data) GetChannel(ch chan data.Data) {
	node2 = <-ch
	append(node2.Node2, "")
	fmt.Println(node2.Node2)

	// send back data to Node1
	ch1 := make(chan data.Data)
	go node2.Run(ch1)
}
