package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	s := <-signals
	switch s {
	case syscall.SIGINT:
		// kill -2 process_id
		fmt.Println("SIGINT")
	case syscall.SIGTERM:
		// kill -15 process_id
		fmt.Println("SIGTERM")
	}
}
