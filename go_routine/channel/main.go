package main

import "fmt"

func routine(ch chan int) {
	fmt.Println(100 + <-ch)
}
func main() {
	channel := make(chan int)

	go routine(channel) //receiver
	//when there is no receiver then leads to deadlock
	channel <- 234 //sender

	fmt.Println("value", channel)
	fmt.Printf("type %T\n", channel)
}

// 2 go routine main and routine func
