package node3

import (
	Data "Golang/src/data"
	"sync"
)

type Node3 struct {
	Channel   chan Data.Data
	sentNodeT map[string]string
}

type INode3 interface {
	Run(node1, t chan Data.Data, wg *sync.WaitGroup)
	GetChannel() chan Data.Data
}

/*
Execute Node3 when receive data
*/
func (n *Node3) Run(node1, T chan Data.Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range n.Channel {

		//fmt.Println("Node33")
		data.Node3 = true
		_, found := n.sentNodeT[data.Id]
		if !found && len(data.Node2) > 0 {
			wg.Add(2)
			//fmt.Println("Node3 -> T")
			if n.sentNodeT == nil {
				n.sentNodeT = make(map[string]string)
			}
			n.sentNodeT[data.Id] = data.Id
			T <- data
			node1 <- data
		}
	}
}

/*
Return Channel for Node3
*/
func (n *Node3) GetChannel() chan Data.Data {
	return n.Channel
}
