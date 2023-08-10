package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started jobs", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
func main() {
	start := time.Now()
	const numjobs = 500
	jobs := make(chan int, numjobs)
	results := make(chan int, numjobs)
	for w := 1; w <= 200; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numjobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numjobs; a++ {
		<-results
	}
	end := time.Now().Sub(start)
	fmt.Println(end)
}
