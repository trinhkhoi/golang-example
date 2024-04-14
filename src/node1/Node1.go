package node1

import (
	Data "Golang/src/data"
	NodeM "Golang/src/m"
	Node2 "Golang/src/node2"
	Node3 "Golang/src/node3"
	"fmt"
	"sync"
)

var ch = make(chan Data.Data)

type INode1 interface {
	Run(node2 chan Data.Data, node3 chan Data.Data, m chan Data.Data, data Data.Data)
	Data.Channel
	ReceiveData(ch chan Data.Data)
}

/*
Processing next action for Node1
*/
func Run(ch2 chan Data.Data, ch3 chan Data.Data, chM chan Data.Data, data Data.Data) {
	var wg sync.WaitGroup
	wg.Add(3)

	// create routine and send data to Node2
	go sendingDataNextNode(&wg, ch2, data)
	Node2.ReceiveData(ch2)

	// create routine and send data to Node3
	go sendingDataNextNode(&wg, ch3, data)

	// create routine and send data to Node M
	go sendingDataNextNode(&wg, chM, data)

	wg.Wait()

}

/*
Return channel for Node1
*/
func GetChannel() chan Data.Data {
	return ch
}

/*
Handling the receive data at Node1
*/
func ReceiveData(chNode1 chan Data.Data) {
	var data = <-chNode1
	data.Node1 = true
	fmt.Println("data.Node1: ", data.Node1)

	var chN2 = Node2.GetChannel()
	var chN3 = Node3.GetChannel()
	var chM = NodeM.GetChannel()
	Run(chN2, chN3, chM, data)
}

/*
Sending data to next channel
*/
func sendingDataNextNode(wg *sync.WaitGroup, ch chan Data.Data, data Data.Data) {
	ch <- data
	defer wg.Done()
}
