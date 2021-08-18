package regex_hu60

import (
	"atlob-spider/core"
	"fmt"
	"regexp"
)

// 论坛板块正则
var forumRe = regexp.MustCompile(`<a href="(user.info.\d+.html)">(\S+)</a>`)
var forumPageRe = regexp.MustCompile(`<li><a href="(bbs.forum.(\S+).html)">\S+</a></li>`)

// ParseForum 把传入的 document 解析为 ParseResult
func ParseForum(document []byte, r core.Request) core.ParseResult {
	var ret core.ParseResult

	forumMatch := forumRe.FindAllSubmatch(document, -1)
	fmt.Printf("\t\t%s 下找到用户 [%d] 个\n", r.Item.Title, len(forumMatch))
	for _, v := range forumMatch {
		ret.Requests = append(ret.Requests, core.Request{
			Url: fmt.Sprintf("%s%s", urlPrefix, v[1]),
			Item: core.Item{
				Title: string(v[2]),
				Parse: core.CreateParse(ParseUser, r),
			},
		})
	}

	forumPageMatch := forumPageRe.FindAllSubmatch(document, -1)
	fmt.Printf("\t%s 下找到论坛 [%d] 个\n", r.Item.Title, len(forumPageMatch))
	for _, v := range forumPageMatch {
		ret.Requests = append(ret.Requests, core.Request{
			Url: fmt.Sprintf("%s%s", urlPrefix, v[1]),
			Item: core.Item{
				Title: string(v[2]),
				Parse: core.CreateParse(ParseForum, r),
			},
		})
	}

	return ret
}
