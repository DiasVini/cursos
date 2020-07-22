package main

import (
	"fmt"
	"runtime"
)

func main() {
	// numero de cpus do pc
	fmt.Println(runtime.NumCPU())
}
