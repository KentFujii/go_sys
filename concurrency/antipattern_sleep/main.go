package main

import (
	"fmt"
	"time"
)

func sub1 (c int) {
	fmt.Println("share by arguments:", c*c)
}

func main(){
	go sub1(10)
	c := 20
	go func() {
		fmt.Println("share by capture", c*c)
	}()
	// 場当たり的にsleepしなければ親スレッドが終了し子スレッドのgoroutineが完了せずに終わってしまう
	time.Sleep(time.Second)
}

