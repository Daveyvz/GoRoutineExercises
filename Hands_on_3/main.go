package main

import (
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var wg sync.WaitGroup
	count := 0
	wg.Add(3)
	var mu sync.Mutex

	go func() {
		mu.Lock()
		v := count
		runtime.Gosched()
		v++
		count = v
		wg.Done()
		mu.Unlock()
	}()

	go func() {
		mu.Lock()
		v := count
		runtime.Gosched()
		v++
		count = v
		wg.Done()
		mu.Unlock()
	}()

	go func() {
		mu.Lock()
		v := count
		runtime.Gosched()
		v++
		count = v
		wg.Done()
		mu.Unlock()
	}()

	wg.Wait()
}
