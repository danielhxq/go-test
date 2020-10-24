package pool

import "sync"

type Pool struct {
	lock      *Lock
	ll        *sync.RWMutex
	pool      []int64
	not_empty *Condition
	not_full  *Condition
}

func init() {

}

func (p *Pool) addElement(newElement int64) {
	for p.isFull() {
		p.not_full.Wait()
	}
	p.pool[0] = newElement
	p.not_empty.Signal()
}

func (p *Pool) removeElement() {
	for p.isEmpty() {
		p.not_empty.Wait()
	}

	p.not_full.Signal()
}

func (p *Pool) isFull() bool {
	p.ll.Lock()
	defer p.ll.Unlock()
	return len(p.pool) == 64
}

func (p *Pool) isEmpty() bool {
	p.ll.Lock()
	defer p.ll.Unlock()
	return len(p.pool) == 0
}
