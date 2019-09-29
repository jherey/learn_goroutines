package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1) // making our channel serve as both the sender and receiver of the message
	ch <- "Hello"

	fmt.Println(<-ch)
}
