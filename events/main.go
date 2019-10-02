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
