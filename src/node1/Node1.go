package node1

import (
	"Golang/src/data"
	"Golang/src/m"
	"Golang/src/node2"
	"Golang/src/node3"
	"fmt"
	"sync"
)

type Data struct {
	data.Data
}

type Node1 interface {
	Run(node2 chan Data, node3 chan Data, m chan Data)
	data.Channel
}

/*
Processing next action for Node1
*/
func (node1 Data) Run(ch2 chan Data, ch3 chan Data, chM chan Data) {
	var wg sync.WaitGroup
	wg.Add(3)

	// create Thread and send data to Node2
	go func() {
		defer wg.Done()
		ch2 <- node1
		node2.Node2.GetChannel(ch2)
	}()

	// create Thread and send data to Node3
	go func() {
		defer wg.Done()
		ch3 <- node1
		node3.Node3.GetChannel(ch3)
	}()

	// create Thread and send data to Node M
	go func() {
		defer wg.Done()
		chM <- node1
		m.NodeM.GetChannel(chM)
	}()

	wg.Wait()

}

/*
Handling when the Node1 receive data
*/
func GetChannel(ch chan Data) {
	node1 := <-ch
	node1.Node1 = true
	fmt.Println(node1.Node1)

	ch2 := make(chan Data)
	ch3 := make(chan Data)
	chM := make(chan Data)

	node1.Run(ch2, ch3, chM)
}
