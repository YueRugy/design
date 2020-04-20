package main

import (
	"fmt"
	"sync"
)

var group sync.WaitGroup

func main() {

	c := make(chan int)
	group.Add(1)
	go sum(c)
	c <- 1
	fmt.Println()
	group.Wait()
}

func sum(c <-chan int) {
	defer group.Done()
	fmt.Println(<-c)
}
