package main

import (
	"crawler1/engine"
	"crawler1/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
		//Url: "https://www.bilibili.com/ranking/all/0/0/1",
		//ParserFunc: bilibiliParser.ParsePaiHang,
	})
	//IPProxy.GetProxy()
}
