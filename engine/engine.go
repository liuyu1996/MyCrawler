package engine

import (
	"crawler1/IPProxy"
	"crawler1/fetcher"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func Run(Seeds ...Request)  {
	var requests []Request

	for _, r := range Seeds{
		requests = append(requests,r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]



		time.Sleep(time.Second * 2)
		log.Printf("Fetching %s",r.Url)
		body, err := fetcher.Fetcher(r.Url)
		if err != nil{
			log.Printf("Fetch:error" + "fetching url %s: %v",r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}
