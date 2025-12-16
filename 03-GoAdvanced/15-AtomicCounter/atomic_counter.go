package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// atomic package is used like mutex but for integers and pointers

type AtomicCounter struct {
	count int64
}

func (a *AtomicCounter) increment() {
	atomic.AddInt64(&a.count, 1)
}

func (a *AtomicCounter) getValue() int64 {
	return a.count
}

func main() {
	var wg sync.WaitGroup
	numOfRoutines := 10
	counter := &AtomicCounter{}
	valueWithoutAtomic := 0

	wg.Add(numOfRoutines)
	for range numOfRoutines {
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				valueWithoutAtomic++
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final incrementer with atomic is: ", counter.getValue())
	fmt.Println("Final incrementer without atomic is: ", valueWithoutAtomic)

}
