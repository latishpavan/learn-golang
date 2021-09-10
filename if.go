package main

import "fmt"

func main() {
	// optional initialization in the if
	if frac := 3 / 2.0; frac > 1 {
		fmt.Println("fraction is bigger")
	}
}
