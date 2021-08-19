package regex_atlob

import (
	"atlob-spider/core"
	"atlob-spider/core/util"
	"fmt"
	"regexp"
)

// PostImgReg 文章图片正则
// 	1 为图片链接
var PostImgReg = regexp.MustCompile("<img src=\"(\\S+)\" alt=\"\" />")

// ParsePostImg 解析文章图片
func ParsePostImg(document []byte, r core.Request) core.ParseResult {
	var ret core.ParseResult

	postImgMatch := PostImgReg.FindAllSubmatch(document, -1)
	fmt.Printf("\t %s 下找到图片 [%d] 个\n", r.Item.Title, len(postImgMatch))
	for _, v := range postImgMatch {
		util.DownloadImg(string(v[1]), r.Item.Title, &ret)
	}

	return ret
}
