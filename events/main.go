package main

import (

)

func main() {

}

// Button struct type
type Button struct {
	eventListeners map[string][]chan string
}

// MakeButton is a button constructor function
func MakeButton() *Button {
	result := new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result
}

// AddEventListener adds an event to a responseChan
func (button *Button) AddEventListener(event string, responseChan chan string) {
	if _, present := button.eventListeners[event]; present {
		button.eventListeners[event] = append(button.eventListeners[event], responseChan)
	} else {
		button.eventListeners[event] = []chan string{responseChan}
	}
}
