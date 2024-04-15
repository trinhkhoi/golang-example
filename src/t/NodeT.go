package node2

import (
	Data "Golang/src/data"
	"fmt"
	"sync"
	"time"
)

type T struct {
	Channel chan Data.Data
}

type NodeT interface {
	Run(wg *sync.WaitGroup)
	GetChannel() chan Data.Data
}

/*
Execute Node T when receive data
*/
func (t *T) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range t.Channel {
		timestamp := time.Now().UTC().Format(time.RFC3339Nano)
		fmt.Printf("T%s-%s-%s\n", data.Id, data.Node2[len(data.Node2)-1], timestamp)
	}

}

/*
Return Channel for NodeT
*/
func (t *T) GetChannel() chan Data.Data {
	return t.Channel
}
