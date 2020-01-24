package main

import (
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	count := 0
	wg.Add(3)
	go func() {
		v := count
		runtime.Gosched()
		v++
		count = v
		wg.Done()
	}()

	go func() {
		v := count
		runtime.Gosched()
		v++
		count = v
		wg.Done()
	}()

	go func() {
		v := count
		runtime.Gosched()
		v++
		count = v
		wg.Done()
	}()

	wg.Wait()
}
