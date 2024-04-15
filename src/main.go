package main

import (
	Data "Golang/src/data"
	NodeM "Golang/src/m"
	Node1 "Golang/src/node1"
	Node2 "Golang/src/node2"
	Node3 "Golang/src/node3"
	NodeT "Golang/src/t"
	"fmt"
	"sync"
)

func ThreadS(N int, node1 chan Data.Data, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < N; i++ {
		data := Data.Data{
			Id:    fmt.Sprintf("%d", i+1),
			Node1: false,
			Node2: []string{},
			Node3: false,
		}
		node1 <- data
	}
}

func main() {
	var N int
	fmt.Print("Enter the limit N: ")
	fmt.Scanln(&N)
	node1 := &Node1.Node1{Channel: make(chan Data.Data, N)}
	node2 := &Node2.Node2{Channel: make(chan Data.Data, N)}
	node3 := &Node3.Node3{Channel: make(chan Data.Data, N)}
	m := &NodeM.M{Channel: make(chan Data.Data, N)}
	t := &NodeT.T{Channel: make(chan Data.Data, N)}

	var wg sync.WaitGroup
	wg.Add(6)
	go ThreadS(N, node1.GetChannel(), &wg)
	go node1.Run(node2.GetChannel(), node3.GetChannel(), m.GetChannel(), &wg)
	go node2.Run(node1.GetChannel(), &wg)
	go node3.Run(node1.GetChannel(), t.GetChannel(), &wg)
	go m.Run(node1.GetChannel(), &wg)
	go t.Run(&wg)
	wg.Wait()

	fmt.Println("Press 'Enter' to exit...")
	fmt.Scanln()

}
