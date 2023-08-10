package main

import (
	"fmt"
	"sync"
)

func routine(i int, wg *sync.WaitGroup) {
	fmt.Println("strated routine", i)
	fmt.Printf("routine %d ended\n", i)
	wg.Done()
}
func main() {
	noOfRoutine := 3
	var wg sync.WaitGroup
	for i := 0; i <= noOfRoutine; i++ {
		wg.Add(1)
		go routine(i, &wg)
	}
	wg.Wait()
	fmt.Println("all routines are finished")

}
