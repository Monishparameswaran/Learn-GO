package main

import (
	"fmt"
	"time"
	"sync"
)


// in this program you can notice the time function ,one goroutine executed before other without the delay it doesnt mean go routine are not executing 
// simultaneously this means when function call one goroutine it is executed in thunder fast before the function enters other function 
// so to see the concurrency use time function
type Counter struct {
	counter int
	sync.Mutex
}

func Increment(done chan bool, obj *Counter) {
	mutex := &obj.Mutex
	for i := 0; i <= 50; i++ {
		fmt.Println(1, " ", i)
		time.Sleep(time.Second * 2)
		mutex.Lock()
		obj.counter++
		mutex.Unlock()
	}
	done <- true
}

func Decrement(done chan bool, obj *Counter) {
	mutex := &obj.Mutex
	for i := 0; i <= 50; i++ {
		fmt.Println(2, " ", i)
		time.Sleep(time.Second * 2)
		mutex.Lock()
		obj.counter--
		mutex.Unlock()
	}
	done <- true
}

func main() {
	done := make(chan bool)
	obj := &Counter{}
	go Increment(done, obj)
	go Decrement(done, obj)
	<-done
	<-done
	fmt.Println(obj.counter)
}
