package node3

import (
	Data "Golang/src/data"
	"fmt"
)

var ch = make(chan Data.Data)

type Node3 interface {
	Run(nodeCh1 chan Data.Data, nodeTCh chan Data.Data)
	Data.Channel
}

func Run(node1Ch chan Data.Data, nodeTCh chan Data.Data) {
	node1Ch <- node3Data
	node3 = <-ch
	node3.Node3 = true
	fmt.Println(node3.Node3)

	// send back data to Node1
	chT := make(chan Data)
	ch <- node3
	node3.Run(ch, chT)
}

func GetChannel() chan Data.Data {
	return ch
}
