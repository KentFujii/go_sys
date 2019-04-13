package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	// 先にgoroutineを用意して後で一斉に実行したい場合はCondを使う
	cond := sync.NewCond(&mutex)
	for _, name := range []string{"A", "B", "C"} {
		go func(name string) {
			mutex.Lock()
			defer mutex.Unlock()
			cond.Wait()
			fmt.Println(name)
		}(name)
	}

	fmt.Println("よーい")
	time.Sleep(time.Second)
	fmt.Println("どん! ")
	cond.Broadcast()
	time.Sleep(time.Second)
}
