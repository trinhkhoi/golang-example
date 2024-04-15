package node2

import (
	Data "Golang/src/data"
	"sync"
	"time"
)

type Node2 struct {
	Channel chan Data.Data
}

type INode2 interface {
	Run(node1 chan Data.Data, wg *sync.WaitGroup)
	GetChannel() chan Data.Data
}

/*
Execute Node2 when receive data
*/
func (n *Node2) Run(node1 chan Data.Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range n.Channel {
		data.Node2 = append(data.Node2, time.Now().UTC().Format(time.RFC3339Nano))
		node1 <- data
	}

}

/*
Return Channel for Node2
*/
func (n *Node2) GetChannel() chan Data.Data {
	return n.Channel
}
