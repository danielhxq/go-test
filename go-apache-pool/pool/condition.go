package pool

import (
	"sync"
)

type Condition struct {
	lock *sync.RWMutex
	ll   *sync.RWMutex
	c    chan int
}

func (c *Condition) Wait() {
	c.lock.Unlock()
	c.ll.Lock()
	c.ll.Unlock()
	defer c.lock.Lock()
	for {
		select {
		case <-c.c:
			break
		}
	}
}

func (c *Condition) Signal() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.c <- 1
}
