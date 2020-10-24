package pool

import "sync"

type Lock struct {
	lock *sync.RWMutex
}
