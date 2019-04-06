package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Fprintf(os.Stderr, "グループID: %d\n", syscall.Getpgrp())
}
