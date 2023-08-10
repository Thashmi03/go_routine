package main

import (
	"fmt"
	"time"
	// "golang.org/x/tools/go/packages"
)

type Item struct {
	ID            int
	Name          string
	PackingEffort time.Duration
}

func PrepareItems(done <-chan bool) <-chan Item {
	items := make(chan Item)
	itemsToShip := []Item{
		{0, "shirt", 1 * time.Second},
		{1, "legos", 1 * time.Second},
		{2, "pant", 1 * time.Second},
		{3, "trouser", 1 * time.Second},
		{4, "shirt", 1 * time.Second},
		{5, "legos", 1 * time.Second},
		{6, "pant", 1 * time.Second},
		{7, "trouser", 1 * time.Second},
		{8, "trouser", 1 * time.Second},
		{9, "shirt", 1 * time.Second},
		{10, "legos", 1 * time.Second},
		{11, "pant", 1 * time.Second},
		{12, "trouser", 1 * time.Second},
	}
	go func() {
		for _, item := range itemsToShip {
			select {
			case <-done:
				return
			case items <- item:
			}
		}
		close(items)
	}()
	return items
}

func PackItems(done <-chan bool, items <-chan Item) <-chan int {
	packages := make(chan int)
	go func() {
		for item := range items {
			select {
			case <-done:
				return
			case packages <- item.ID:
				time.Sleep(item.PackingEffort)
			}
		}
		close(packages)
	}()
	return packages
}

func main() {
	done := make(chan bool)
	defer close(done)
	start := time.Now()
	packages := PackItems(done, PrepareItems(done))
	numPackages := 0
	for p := range packages {
		numPackages++
		fmt.Printf("shipping packages no.%d\n", p)
	}
	fmt.Printf("took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)
}
