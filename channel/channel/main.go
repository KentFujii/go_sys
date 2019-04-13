package main

import (
	"fmt"
)

func main() {
	fmt.Println("start sub()")
	// チャネルとは、並列アクセスからでも順序を保証できるキューのようなもの
	// 終了を受け取るためのチャネル
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is finished")
		// 終了を通知
		done <- true
	}()
	// 終了を待つ
	<-done
	fmt.Println("all tasks are finished")
}
