package parser

import (
	"crawler1/engine"
	"fmt"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">`)

func ParseCity(contents []byte)  engine.ParseResult{
	matches := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	if matches == nil {
		fmt.Println("matches is nil")
	}

	for _, m := range matches{
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:         string(m[1]),
			ParserFunc : ParseProfile,
		})
		//fmt.Printf("City:%s, url:%s\n", m[2], m[1])
	}
	return result
}
