package regex_hu60

import (
	"atlob-spider/core"
	"fmt"
	"regexp"
)

var uidRe = regexp.MustCompile(`<td>UID：</td>\s+<td>(\d+)</td>`)

// var nameRe = regexp.MustCompile(`<td>用户名：</td>\s+<td>(\S+)</td>`)
var signatureRe = regexp.MustCompile(`<td>个性签名：</td>\s+<td>(\S+)</td>`)
var contactRe = regexp.MustCompile(`<td>联系方式：</td>\s+<td>([^<]+)</td>`)
var regTimeRe = regexp.MustCompile(`<td>注册时间：</td>\s+<td>([^<]+)</td>`)
var avatarSrcRe = regexp.MustCompile(`<img src="([^\"]+)" width="96"/><br/>`)

// ParseUser 把传入的 document 解析为 ParseResult
func ParseUser(document []byte, r core.Request) core.ParseResult {
	var ret core.ParseResult

	// 分类列表正则匹配,将匹配结果装入对象 ParseResult
	postImgMatch := avatarSrcRe.FindAllSubmatch(document, -1)
	fmt.Printf("\t\t\t%s 下找到头像 [%d] 个\n", r.Item.Title, len(postImgMatch))
	for i, v := range postImgMatch {
		ret.Requests = append(ret.Requests, core.Request{
			Url: string(v[1]),
			Item: core.Item{
				Title: fmt.Sprintf("%s-头像-%d", r.Item.Title, i),
				Parse: core.CreateParse(core.NilParser, r),
			},
		})
	}

	return ret
}
