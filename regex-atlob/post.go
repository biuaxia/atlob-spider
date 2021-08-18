package regex_atlob

import (
	"atlob-spider/core"
	"fmt"
	"regexp"
)

// PostImgReg 文章图片正则
// 	1 为图片链接
var PostImgReg = regexp.MustCompile("<img src=\"(\\S+)\" alt=\"\" />")

// ParsePostImg 解析文章图片
func ParsePostImg(document []byte, r core.Request) core.ParseResult {
	var ret core.ParseResult

	// 分类列表正则匹配,将匹配结果装入对象 ParseResult
	postImgMatch := PostImgReg.FindAllSubmatch(document, -1)
	fmt.Printf("\t %s 下找到图片 [%d] 个\n", r.Item.Title, len(postImgMatch))
	for i, v := range postImgMatch {
		ret.Requests = append(ret.Requests, core.Request{
			Url: string(v[1]),
			Item: core.Item{
				Title: fmt.Sprintf("%s-图片-%d", r.Item.Title, i),
				Parse: core.CreateParse(core.NilParser, r),
			},
		})
	}

	return ret
}
