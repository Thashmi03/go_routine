package main

import (
	"fmt"
	"time"
)
func num(x int) {
	fmt.Println(x)
}
func main() {
	for j := 0; j < 20; j++ {
		go num(j)
		// time.Sleep(1 * time.Second)
	}
	time.Sleep(1 * time.Second)
}
