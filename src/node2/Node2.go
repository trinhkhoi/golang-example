package node2

import (
	Data "Golang/src/data"
	Node1 "Golang/src/node1"
	"fmt"
	"time"
)

var ch = make(chan Data.Data)

type Node2 interface {
	Run(ch1 chan Data.Data)
	Data.Channel
}

/*
Processing next action for Node2
*/
func Run(ch1 chan Data.Data) {
	Node1.ReceiveData(ch1)
}

/*
Return channel for Node2
*/
func GetChannel() chan Data.Data {
	return ch
}

/*
Handling the received data at Node2
*/
func ReceiveData(ch2 chan Data.Data) {
	var data = <-ch2
	t := time.Now()
	arrNode2s := append(data.Node2, t.String())
	data.Node2 = arrNode2s
	fmt.Println("data.Node2", data.Node2)

	// send back data to Node 1
	var ch1 = Node1.GetChannel()
	go func() {
		ch1 <- data
	}()
	Run(ch1)
}
