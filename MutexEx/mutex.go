package mutexPac

import (
	"fmt"
	"sync"
)

var counter int
var counterMutex sync.Mutex

func MutexEx() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Increment()
		}()
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

func Increment() {
	counterMutex.Lock()
	defer counterMutex.Unlock()
	counter++
}
