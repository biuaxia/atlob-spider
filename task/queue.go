package task

import (
	"atlob-spider/core"
	"runtime/pprof"
	"time"
)

type Queued struct {
	requestChan chan core.Request
	workerChan  chan chan core.Request
}

func (q *Queued) Submit(request core.Request) {
	q.requestChan <- request
}

func (q *Queued) WorkChan() chan core.Request {
	return make(chan core.Request)
}

func (q *Queued) WorkerReady(worker chan core.Request) {
	q.workerChan <- worker
}

var threadProfile = pprof.Lookup("threadcreate")

func (q *Queued) Run() {
	// 准备任务调度的channel
	q.requestChan = make(chan core.Request)
	q.workerChan = make(chan chan core.Request)

	// 开始循环调度通信
	go func() {
		// 请求队列
		var requestQueue []core.Request
		// 工作队列
		var workerQueue []chan core.Request

		for {
			var activeRequest core.Request
			var activeWorker chan core.Request

			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}

			// 调度
			select {
			case r := <-q.requestChan:
				time.Sleep(200 * time.Millisecond)
				requestQueue = append(requestQueue, r)
			case r := <-q.workerChan:
				workerQueue = append(workerQueue, r)
			case activeWorker <- activeRequest:
				// 请求分析完毕交给工作channel进行实际处理
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}

		}
	}()

}
