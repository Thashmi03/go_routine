package main

import (
	"fmt"
	"time"
)

func hi() {
	fmt.Println("routine") //so this is not printed
	//once main complete then no thread allocation so it end and not printed
}
func main() {
	fmt.Println("main start")
	go hi()
	// for giving time for routine we use sleep for main
	time.Sleep(10 * time.Microsecond)
	fmt.Println("main end")
}

//not providing enough time to execute for go routine
