package main

import (
	"fmt"
	"time"
)

//without go routine

// func fib(n1,n2 ,n int){
// 	if(n==0){
// 		return
// 	}
// 	n3:=n1+n2
// 	//fmt.Println(n3)
// 	fib(n2,n3,n-1)

// }

// func main(){
// 	start:=time.Now()
// 	fib(0,1,50)
// 	end:=time.Now().Sub(start)
// 	fmt.Println(end)
// }

// with routine

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i <= n; i++ {
		c <- x //sender
		x, y = y, x+y
	}
	close(c)
}
func main() {
	start := time.Now()
	c := make(chan int, 50)
	go fib(cap(c), c) //receiver
	end := time.Now().Sub(start)
	fmt.Println(end)
}
