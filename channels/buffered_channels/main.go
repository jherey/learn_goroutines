package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "These are the time's that try men's soul\n"

	words := strings.Split(phrase, "")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	// for {
	// 	if msg, ok := <- ch; ok {
	// 		fmt.Print(msg + " ")
	// 	} else {
	// 		break
	// 	}
	// }

	//	OR
	for msg := range ch {
		fmt.Print(msg + " ")
	}

	// ch <- "test"	// we can't pass a msg into the channel anymore as it's closed.
}
