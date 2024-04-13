package node2

import (
	"Golang/src/data"
	"Golang/src/node1"
	"fmt"
	"time"
)

type Data struct {
	data.Data
}

type Node2 interface {
	Run(node1 chan Data)
	data.Channel
}

func (data Data) Run(node1Ch chan Data) {
	node1Ch <- data
	// send back data to Node1
	node1.GetChannel(node1Ch)
}

func (node2 Data) GetChannel(ch chan Data) {
	node2 = <-ch
	t := time.Now()
	arrNode2s := append(node2.Node2, t.String())
	node2.Node2 = arrNode2s
	fmt.Println(node2.Node2)

	// call Run function to execute the next step
	ch1 := make(chan Data)
	node2.Run(ch1)
}
