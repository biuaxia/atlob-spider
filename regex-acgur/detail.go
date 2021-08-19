package regex_acgur

import (
	"atlob-spider/core"
	"atlob-spider/core/util"
	"fmt"
	"regexp"
)

const FileStorePath = "./tmp/"

var ImgReg = regexp.MustCompile(`<img crossorigin data-src="([^\"]+)">`)

func ParseDetailJson(document []byte, r core.Request) core.ParseResult {
	var ret core.ParseResult

	categoryMatch := ImgReg.FindAllSubmatch(document, -1)
	fmt.Printf("\t %s 下找到图片 [%d] 个\n", r.Item.Title, len(categoryMatch))
	for _, v := range categoryMatch {
		util.DownloadImg(string(v[1]), r.Item.Title, &ret)
	}

	return ret
}
