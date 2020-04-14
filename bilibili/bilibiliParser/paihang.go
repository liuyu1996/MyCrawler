package bilibiliParser

import (
	"crawler1/engine"
	"crawler1/model"
	"fmt"
	"regexp"
	"strconv"
)

var paihangRe = regexp.MustCompile(`<a href="(https://www.bilibili.com/video/[0-9a-zA-Z]+)" target="_blank" class="title">([^<]+)</a>`)
var playCountRe = regexp.MustCompile(`<span class="data-box"><i class="b-icon play"></i>([^<]+)</span>`)
var danmuCountRe = regexp.MustCompile(`<span class="data-box"><i class="b-icon view"></i>([^<]+)</span>`)
var scoresRe = regexp.MustCompile(`<div class="pts"><div>([0-9]+)</div>`)

func ParsePaiHang(contents []byte)  engine.ParseResult{
	matches := paihangRe.FindAllSubmatch(contents, -1)
	videoinfo  := model.VideoInfo{}
	scores, err := strconv.Atoi(extractString(contents, scoresRe))
	if err == nil{
		videoinfo.Scores = scores
	}
	videoinfo.DanMuCount = extractString(contents, danmuCountRe)
	videoinfo.PlayCount = extractString(contents, playCountRe)
	result := engine.ParseResult{}
	limit := 5
	if matches == nil {
		fmt.Println("matches is nil")
	}

	for _, m := range matches{
		title := string(m[2])
		videoinfo.Title = title
		result.Items = append(result.Items, "title:" + title)
		result.Requests = append(result.Requests, engine.Request{
			Url:         string(m[1]),
			ParserFunc : func(c []byte) engine.ParseResult {
				return parseBilibiliProfile(c, videoinfo)
			},
		})
		limit--
		if limit == 0{
			break
		}
		//fmt.Printf("City:%s, url:%s\n", m[2], m[1])
	}
	return result
}
