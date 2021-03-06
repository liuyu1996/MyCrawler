package parser

import (
	"crawler1/engine"
	"fmt"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte)  engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 5
	if matches == nil {
		fmt.Println("matches is nil")
	}

	for _, m := range matches{
		result.Items = append(result.Items, "City" + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:         string(m[1]),
			ParserFunc : ParseCity,
		})
		limit--
		if limit == 0{
			break
		}
		//fmt.Printf("City:%s, url:%s\n", m[2], m[1])
	}
	return result
}