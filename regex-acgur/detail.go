package regex_acgur

import (
	"atlob-spider/core"
	"bufio"
	bytes2 "bytes"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
)

const FileStorePath = "./tmp/"

var ImgReg = regexp.MustCompile(`<img crossorigin data-src="([^\"]+)">`)

func ParseDetailJson(document []byte, r core.Request) core.ParseResult {
	var ret core.ParseResult

	categoryMatch := ImgReg.FindAllSubmatch(document, -1)
	fmt.Printf("\t %s 下找到图片 [%d] 个\n", r.Item.Title, len(categoryMatch))
	for _, v := range categoryMatch {
		filename := path.Base(string(v[1]))
		ret.Requests = append(ret.Requests, core.Request{
			Url: string(v[1]),
			Item: core.Item{
				Title: r.Item.Title,
				Parse: core.CreateParse(func(bytes []byte, r core.Request) core.ParseResult {
					imgPath := fmt.Sprintf("%s%s/", FileStorePath, r.Item.Title)
					if !isExist(imgPath) {
						// 递归创建文件夹
						err := os.MkdirAll(imgPath, os.ModePerm)
						if err != nil {
							fmt.Printf("创建文件夹出错: %v\n", err)
							return core.ParseResult{}
						}
					}

					reader := bufio.NewReaderSize(bytes2.NewReader(bytes), 32*1024)

					// 最终文件存储路径
					finalPath := imgPath + filename

					flag := isExist(finalPath)
					// 如果存在则判断文件大小是否为空，不为空则跳过写入，为空或者文件不存在都走入下一步
					if flag {
						stat, err := os.Stat(finalPath)
						if nil != err {
							fmt.Printf("获取文件状态出错: %v\n", err)
							flag = true
						}
						if stat.Size() == 0 {
							flag = true
						}
					}

					var written int64

					// 创建文件
					if !flag {
						file, err := os.Create(finalPath)
						if err != nil {
							fmt.Printf("创建文件夹出错: %v\n", err)
							return core.ParseResult{}
						}
						// 获得文件的writer对象
						writer := bufio.NewWriter(file)

						written, _ = io.Copy(writer, reader)
					}
					fmt.Printf("处理文件完成 -> %q | %-10dbytes\n", filename, written)
					return core.ParseResult{}
				}, core.Request{}),
			},
		})
	}

	return ret
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
