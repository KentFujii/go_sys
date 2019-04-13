package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// シグナルを受けるチャネルを用意
	signals := make(chan os.Signal, 1)
	// SIGINTを受け取る
	signal.Notify(signals, syscall.SIGINT)
	fmt.Println("Waiting SIGINT (CTRL+C)")
	// OSからのシグナルをチャネルで受け取る
	<-signals
	fmt.Println("SIGINT arrived")
}
