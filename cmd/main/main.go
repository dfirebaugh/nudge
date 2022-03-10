package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	println()
	println("running nudge... ")
	println("[close this window to exit]")
	schedule(time.Minute*29, nudge)
	fmt.Scanln() // don't exit
}

func nudge() {
	println("nudge ----- ", time.Now().String())
	robotgo.MoveRelative(-1, 0)
}

func schedule(interval time.Duration, tickFunc func()) {
	ticker := time.NewTicker(interval)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				go tickFunc()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
