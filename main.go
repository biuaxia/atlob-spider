package main

import (
	"atlob-spider/core"
	"atlob-spider/engine"
	regex_acgur "atlob-spider/regex-acgur"
	"atlob-spider/task"
)

// main 项目入口
func main() {
	r := core.Request{
		Url: "https://api.acgur.com/ranking/daily/all/1?utm_source=biuaxia.cn",
		Item: core.Item{
			Parse: core.CreateParse(regex_acgur.ParseIndexJson, core.Request{}),
			Title: "种子页面",
		},
	}

	concurrentEngine := engine.ConcurrentEngine{
		Task:      &task.Queued{},
		WorkCount: 1024,
	}

	concurrentEngine.Run(r)
}
