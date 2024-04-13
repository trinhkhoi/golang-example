package main

import (
	"Golang/src/data"
	"Golang/src/node1"
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch := make(chan data.Data)
	fmt.Println("begin")
	for i := 0; i < 10; i++ {
		data := data.Data{strconv.FormatInt(time.Now().UnixNano(), 10), false, []string{}, false}
		ch <- data
		go node1.Node1.GetChannel(ch)
	}
	time.Sleep(time.Second)
}
