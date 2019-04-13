package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	// 同時に実行されると困る処理を一つに制限する
	// 下の例はidのインクリメントが同時に一つしか実行されないように制限している
	mutex.Lock()
	id++
	defer mutex.Unlock()
	return id
}
func main() {
	var  mutex sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
		}()
	}
}
