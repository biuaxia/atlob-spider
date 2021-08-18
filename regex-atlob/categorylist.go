package regex_atlob

import (
	"atlob-spider/core"
	"fmt"
	"regexp"
)

// CategoryListReg 分类列表匹配正则
// 	1 为链接
// 	2 为主题
var CategoryListReg = regexp.MustCompile("<li><a href=\"(\\S+)\" class=\"\\S{0,2}\" rel=\"category\">(\\S+)</a></li>")

// ParseCategoryList 解析分类列表
func ParseCategoryList(document []byte, r core.Request) core.ParseResult {
	var ret core.ParseResult

	// 分类列表正则匹配,将匹配结果装入对象 ParseResult
	categoryListMatch := CategoryListReg.FindAllSubmatch(document, -1)
	fmt.Printf("%s 下找到分类 [%d] 个\n", r.Item.Title, len(categoryListMatch))
	for _, v := range categoryListMatch {
		ret.Requests = append(ret.Requests, core.Request{
			Url: string(v[1]),
			Item: core.Item{
				Title: string(v[2]),
				Parse: core.CreateParse(ParseCategory, r),
			},
		})
	}

	return ret
}
