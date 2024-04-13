package node3

import (
	"Golang/src/data"
	"fmt"
	"time"
)

type Data struct {
	data.Data
}
type Node3 interface {
	Run(node1 chan Data)
	data.Channel
}

func (data Data) Run(node1Ch chan Data) {
	node1Ch <- data
}

func (node2 Data) GetChannel(ch chan Data) {
	node2 = <-ch
	t := time.Now()
	append(node2.Node2, t.String())
	fmt.Println(node2.Node2)

	// send back data to Node1
	ch1 := make(chan Data)
	go node2.Run(ch1)
}
