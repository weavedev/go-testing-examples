package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	delayedFunc("this is a function which prints a message with a small delay")
}

// delayedFunc a function which prints a message with some build in delay
func delayedFunc(message string) {
	for i := 0; i < 10; i++ {
		log.Info(message)
		time.Sleep(time.Millisecond * 100)
	}
}
