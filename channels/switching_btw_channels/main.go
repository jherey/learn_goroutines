package main

import (
	"fmt"
)

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	/*
	msg := Message{
		To: []string{"olufayojeremiah@gmailcom"},
		From: "jolufayo@gmail.com",
		Content: "Keep it secret, keep it safe",
	}

	failedMessage := FailedMessage{
		ErrorMessage: "Message intercepted by black rider",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage
	*/

	select {
	case receivedMsg := <- msgCh:
		fmt.Println(receivedMsg)
	case receivedErr := <- errCh:
		fmt.Println(receivedErr)
	default:
		fmt.Println("No message received")
	}
}

// Message data type
type Message struct {
	To []string
	From string
	Content string
}

// FailedMessage data type
type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}
