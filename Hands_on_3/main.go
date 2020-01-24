package main

import (
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var wg sync.WaitGroup
	var count int64
	wg.Add(3)
	//var mu sync.Mutex

	go func() {
		//	mu.Lock()
		/*v := count
		runtime.Gosched()
		v++
		count = v */
		wg.Done()
		//	mu.Unlock()
		atomic.AddInt64(&count, 1)

	}()

	go func() {
		//	mu.Lock()
		/*v := count
		runtime.Gosched()
		v++
		count = v */
		wg.Done()
		//	mu.Unlock()
		atomic.AddInt64(&count, 1)
	}()

	go func() {
		//	mu.Lock()
		/*v := count
		runtime.Gosched()
		v++
		count = v */
		wg.Done()
		//	mu.Unlock()
		atomic.AddInt64(&count, 1)
	}()

	wg.Wait()
}
