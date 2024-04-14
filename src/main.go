package main

import (
	Data "Golang/src/data"
	Node1 "Golang/src/node1"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("begin")
	var ch = Node1.GetChannel()
	for i := 0; i < 10; i++ {
		data := Data.Data{strconv.FormatInt(time.Now().UnixNano(), 10), false, []string{}, false}
		go func() {
			ch <- data
		}()
		Node1.ReceiveData(ch)
	}
	time.Sleep(time.Second)
}
