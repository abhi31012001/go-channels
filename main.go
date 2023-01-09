package main

import (
	"fmt"
	"sync"
)

func main() {
	cha := make(chan int, 1)
	wc := &sync.WaitGroup{}
	wc.Add(2)
	go func(ch <-chan int, wc *sync.WaitGroup) {
		c, er := <-cha
		fmt.Println(er)
		fmt.Println(c)
		wc.Done()
	}(cha, wc)

	go func(ch <-chan int, wc *sync.WaitGroup) {
		cha <- 5
		close(cha)
		wc.Done()
	}(cha, wc)
	wc.Wait()
}
