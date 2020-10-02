package main

import (
	"awesomeProject3/crawler/engine"
	"awesomeProject3/crawler/scheduler"
	"awesomeProject3/crawler/zhenai/parser"
)

func main() {
	request := engine.Request{Url: "http://www.zhenai.com/zhenghun", ParserFunc: parser.ParserCityList}
	concurrentEngine := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	concurrentEngine.Run(request)
}
