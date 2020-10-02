package scheduler

import "awesomeProject3/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q *QueuedScheduler) ConfigureWorkerChan(requests chan engine.Request) {
	panic("implement me")
}

func (q *QueuedScheduler) WorkReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) Run() {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-q.requestChan:
				// send r to a ? work
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				// send ?next_request to w?
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
