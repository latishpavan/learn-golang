package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}

		close(ch)
	}()

	for val := range ch {
		fmt.Printf("recieved %d\n", val)
	}
}
