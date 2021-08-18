package regex_hu60

import (
	"atlob-spider/core"
	"fmt"
	"regexp"
)

const urlPrefix = "https://hu60.cn/q.php/"

const avatarUrlPrefix = "https://hu60.cn/"

// 论坛板块列表正则
var forumListRe = regexp.MustCompile(`<a href="(bbs.forum.\d+.html)"\s>(\S*)</a>`)

// ParseForumList 把传入的 document 解析为 ParseResult
func ParseForumList(document []byte, r core.Request) core.ParseResult {
	fmt.Println(string(document))
	var ret core.ParseResult

	// 分类列表正则匹配,将匹配结果装入对象 ParseResult
	categoryListMatch := forumListRe.FindAllSubmatch(document, -1)
	fmt.Printf("\t%s 下找到论坛 [%d] 个\n", r.Item.Title, len(categoryListMatch))
	for _, v := range categoryListMatch {
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
