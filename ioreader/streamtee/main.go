package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	// teeとは英語のTであり、T字路を思い描くとその機能がイメージしやすい
	teeReader := io.TeeReader(reader, &buffer)
	// データを捨てる
	_, _ = ioutil.ReadAll(teeReader)
	// けどバッファーに残ってる
	fmt.Println(buffer.String())
}
