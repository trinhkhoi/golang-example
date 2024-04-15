package m

import (
	Data "Golang/src/data"
	"fmt"
	"sync"
	"time"
)

type MData struct {
	Id        string
	LastNode2 string
	timestamp string
}

type M struct {
	Channel   chan Data.Data
	storeLock sync.Mutex
	store     map[string][]MData
}

type NodeM interface {
	Run(node1 chan Data.Data, wg *sync.WaitGroup)
	GetChannel() chan Data.Data
}

/*
Execute Node M when receive data
*/
func (m *M) Run(node1 chan Data.Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range m.Channel {
		_, found := m.store[data.Id]
		if !found && m.store == nil {
			m.store = map[string][]MData{
				data.Id: {},
			}
		} else if !found {
			m.store[data.Id] = []MData{}
		}
		//fmt.Printf("NodeM11: %d-%d\n", len(m.store[data.Id]), len(data.Node2))
		if len(m.store[data.Id]) < len(data.Node2) {
			wg.Add(1)
			lastNode2data := data.Node2[len(data.Node2)-1]
			timestamp := time.Now().UTC().Format(time.RFC3339Nano)
			m.storeLock.Lock()
			mData := m.store[data.Id]
			mData = append(mData, MData{
				Id:        data.Id,
				LastNode2: lastNode2data,
				timestamp: timestamp,
			})
			m.store[data.Id] = mData
			m.storeLock.Unlock()
			fmt.Printf("%s - %s - %s\n", data.Id, lastNode2data, timestamp)
			node1 <- data
		}
	}
}

/*
Return Channel for NodeM
*/
func (m *M) GetChannel() chan Data.Data {
	return m.Channel
}
