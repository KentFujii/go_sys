package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		go func() {
			// 子スレッドであるgoroutineが起動するときには親スレッドのループが回りきってしまう
			// よって全部のtaskが最後のタスクになってしまう
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}
