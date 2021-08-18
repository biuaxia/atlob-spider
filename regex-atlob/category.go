package regex_atlob

import (
	"atlob-spider/core"
	"fmt"
	"regexp"
)

// CategoryReg 分类匹配正则
// 	1 为链接
// 	2 为主题
var CategoryReg = regexp.MustCompile("<a href=\"(\\S+)\" title=\"([^\"]+)\" class=\"fancyimg home-blog-entry-thumb\" rel=\"\\S*\">")

// ParseCategory 解析分类
func ParseCategory(document []byte, r core.Request) core.ParseResult {
	var ret core.ParseResult

	categoryMatch := CategoryReg.FindAllSubmatch(document, -1)
	fmt.Printf("\t %s 下找到分类 [%d] 个\n", r.Item.Title, len(categoryMatch))
	for _, v := range categoryMatch {
		ret.Requests = append(ret.Requests, core.Request{
			Url: string(v[1]),
			Item: core.Item{
				Title: string(v[2]),
				Parse: core.CreateParse(ParsePostImg, r),
			},
		})
	}

	return ret
}
