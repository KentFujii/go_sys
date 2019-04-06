package main

import (
	"fmt"
	"os"
)

func main () {
	fmt.Printf("ユーザID: %d\n", os.Getuid())
	fmt.Printf("グループID: %d\n", os.Getgid())
	fmt.Printf("実効ユーザID: %d\n", os.Geteuid())
	fmt.Printf("実効グループID: %d\n", os.Getegid())
}

