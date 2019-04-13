package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("start sub()")
	// 終了を受け取るコンテキストctxと、そのコンテキストを終了させるcancel関数
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}
