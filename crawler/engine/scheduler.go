package engine

type SimpleScheduler struct {
	workerChan chan Request
}

func (simpleScheduler *SimpleScheduler) Submit(request Request) {
	go func() {
		simpleScheduler.workerChan <- request
	}()
}

func (simpleScheduler *SimpleScheduler) ConfigMasterWorkerChan(in chan Request) {
	simpleScheduler.workerChan = in
}
