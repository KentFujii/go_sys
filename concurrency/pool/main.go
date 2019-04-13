package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	// Poolはオブジェクトのキャッシュを実現する仕組み
	pool := sync.Pool {
		New: func() interface {} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}
	pool.Put("manualy added: 1")
	pool.Put("manualy added: 2")
	// 上の追加したキャッシュからデータを受け取れる
	// ただしGCが走るとキャッシュが消えてしまうので注意すること
	// GCの実行は runtime.GC() で出来る
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	// 下の処理だけ新規作成
	fmt.Println(pool.Get())
}
