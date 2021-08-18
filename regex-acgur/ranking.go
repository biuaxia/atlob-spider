package regex_acgur

import (
	"atlob-spider/core"
	"encoding/json"
	"fmt"
)

// ParseIndexJson 首页Json解析
func ParseIndexJson(document []byte, r core.Request) core.ParseResult {
	var ret core.ParseResult

	var rankingAll RankingAll
	_ = json.Unmarshal(document, &rankingAll)

	fmt.Printf("\t%s 下解析到json对象 [%d] 个\n", r.Item.Title, len(rankingAll.Data))

	for _, j := range rankingAll.Data {
		ret.Requests = append(ret.Requests, core.Request{
			Url: fmt.Sprintf("https://acgur.com/detail/%d", j.IllustId),
			Item: core.Item{
				Title: j.Title,
				Parse: core.CreateParse(ParseDetailJson, r),
			},
		})
	}

	val, ok := rankingAll.Next.(float64)
	i := int(val)
	if ok && rankingAll.Page < i{
		ret.Requests = append(ret.Requests, core.Request{
			Url: fmt.Sprintf("https://api.acgur.com/ranking/daily/all/%d?utm_source=biuaxia.cn", i),
			Item: core.Item{
				Parse: core.CreateParse(ParseIndexJson, core.Request{}),
				Title: fmt.Sprintf("第 %d 页", i),
			},
		})
	}

	return ret
}

type RankingAll struct {
	Title string `json:"title"`
	Mode  string `json:"mode"`
	Type  string `json:"type"`
	Data  []struct {
		Id              int    `json:"id"`
		IllustId        int    `json:"illust_id"`
		UserId          int    `json:"user_id"`
		UserName        string `json:"user_name"`
		Title           string `json:"title"`
		Date            string `json:"date"`
		Width           int    `json:"width"`
		Height          int    `json:"height"`
		Rank            int    `json:"rank"`
		YesRank         int    `json:"yes_rank"`
		IllustPageCount string `json:"illust_page_count"`
		Attr            string `json:"attr"`
		Url             string `json:"url"`
		ProfileImg      string `json:"profile_img"`
	} `json:"data"`
	Page     int         `json:"page"`
	Prev     interface{} `json:"prev"`
	Next     interface{} `json:"next"`
	Date     string      `json:"date"`
	PrevDate string      `json:"prev_date"`
	NextDate bool        `json:"next_date"`
}
