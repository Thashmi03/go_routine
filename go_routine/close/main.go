package main

import "fmt"

func myfunc(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "go lang is awesome"
	}
	close(c)
}
func main() {
	// main act as receiver so no need of wait statement
	c := make(chan string, 8)
	go myfunc(c)
	count := 0
	for {
		// acting as receiver and verifier
		res, ok := <-c // ok says whether it is open (true) or not(false)
		count++
		if !ok {
			fmt.Println("close", ok)
			break
		}
		// fmt.Printf("%v %T",res,ok)
		fmt.Println("channel is open", count, res, ok)
	}
}
