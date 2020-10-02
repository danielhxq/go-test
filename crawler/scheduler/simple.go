package scheduler

import "awesomeProject3/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker channel
	go func() { s.workerChan <- request }()
}

func (s *SimpleScheduler) ConfigureWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
