package node1

import (
	Data "Golang/src/data"
	"sync"
)

type Node1 struct {
	Channel chan Data.Data
}

type INode1 interface {
	Run(node2, node3, m chan Data.Data, wg *sync.WaitGroup)
	GetChannel() chan Data.Data
}

func (n *Node1) Run(node2, node3, m chan Data.Data, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println("Channel1: ", len(n.Channel))
	for data := range n.Channel {
		//fmt.Println("Node11")
		if !data.Node1 {
			//fmt.Println("Node1")
			data.Node1 = true
			node2 <- data
		}
		if len(data.Node2) > 0 {
			node3 <- data
			m <- data
		}
	}
}

/*
Return channel for Node1
*/
func (n *Node1) GetChannel() chan Data.Data {
	return n.Channel
}
