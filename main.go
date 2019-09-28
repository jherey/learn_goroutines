package main

import (
	"time"
	"runtime"
)

func main() {

	godur, _ := time.ParseDuration("10ms")
	runtime.GOMAXPROCS(2) // allowing the code to run with 2 processor threads, this brings about uncertain behaviours as any thread can take over at anytime

	go func() {
		for i := 0; i < 100; i++ {
			println("Hello")
			time.Sleep(godur)
		}
	}() // anonymous function

	go func() {
		for i := 0; i < 100; i++ {
			println("Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
