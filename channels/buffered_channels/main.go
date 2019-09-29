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

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}

	// ch <- "test"	// we can't pass a msg into the channel anymore as it's closed.
}
