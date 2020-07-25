package main
//Mutex: Data Structure for Mutual Exclusion
//Go SDL for it is sync.Mutex and its two methods are:
//Lock
//Unlock
import (
	"fmt"
	"sync"
)
type AtomicInt struct {
	mu sync.Mutex
	n int
}
func (a *AtomicInt) Add(n int) {
	a.mu.Lock()
	a.n+=n
	a.mu.Unlock()
}

func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}

func Demo_Mutex() {
	wait := make(chan struct{})
	var n AtomicInt
	go func() {
		n.Add(1)
		close(wait)
	}()
	n.Add(1)
	<-wait
	fmt.Println("To demo MutEx",n.Value())
}