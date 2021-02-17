package engine

import "log"

type CrawlerEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)
	ConfigMasterWorkerChan(chan Request)
}

func (e *CrawlerEngine) Run(seeds ...Request) {
	in := make(chan Request)      //scheduler 的输入
	out := make(chan ParseResult) //worker的输出
	e.Scheduler.ConfigMasterWorkerChan(in)

	//创建goroutine
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	//把任务提交给scheduler
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("解析结果(%d): %v\n", itemCount, item)
			itemCount++
		}

		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}
	}
}

//创建任务,调用worker,分发goroutine
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
