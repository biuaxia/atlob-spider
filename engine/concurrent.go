package engine

import (
	"atlob-spider/core"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ConcurrentEngine struct {
	core.Task
	WorkCount int
}

func (c ConcurrentEngine) Run(seeds ...core.Request) {
	resp := make(chan core.ParseResult)

	// 任务调度前的准备goon工作
	c.Task.Run()

	for i := 0; i < c.WorkCount; i++ {
		// 创建工作进程
		createWorker(c.Task.WorkChan(), resp, c.Task)
	}

	// 任务提交
	for _, i := range seeds {
		if isDuplicate(i.Url) {
			continue
		}
		c.Task.Submit(i)
	}

	for {
		result := <-resp
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			c.Task.Submit(request)
		}
	}

}

func createWorker(in chan core.Request, resp chan core.ParseResult, notifier core.ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			resp <- result
		}
	}()
}

func worker(r core.Request) (core.ParseResult, error) {
	url := r.Url

	fmt.Printf("处理任务 %q %q 中\n", r.Url, r.Item.Title)

	response, err := http.Get(r.Url)
	if err != nil || http.StatusOK != response.StatusCode {
		fmt.Errorf("Exec failed -> [%s: %s]\n", r.Item.Title, r.Url)
		return core.ParseResult{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		fmt.Printf("Fetcher failed, Info ->  %s: %s\n", url, err)
		return core.ParseResult{}, err
	}

	fmt.Printf("处理任务 %q %q 完毕\n", r.Url, r.Item.Title)
	return r.Item.Parse.Parser(body, r), nil
}

var urlList = make(map[string]bool)

func isDuplicate(url string) bool {
	if urlList[url] {
		return true
	}
	urlList[url] = true
	return false
}
