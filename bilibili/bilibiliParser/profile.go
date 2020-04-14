package bilibiliParser

import (
	"crawler1/engine"
	"crawler1/model"
	"regexp"
	"strconv"
)

var authorRe = regexp.MustCompile(`<meta data-vue-meta="true" itemprop="author" name="author" content="([^>]+)">`)
var commentCountRe = regexp.MustCompile(`<meta data-vue-meta="true" itemprop="commentCount" content="([0-9]+)">`)
var datePublishedRe = regexp.MustCompile(`<meta data-vue-meta="true" itemprop="datePublished" content="([^>]+)">`)

func parseBilibiliProfile(contents []byte, v model.VideoInfo)  engine.ParseResult{
	profile := model.BilibiliProfile{}

	profile.Title = v.Title
	profile.PlayCount = v.PlayCount
	profile.DanMuCount = v.DanMuCount
	commentCount, err := strconv.Atoi(extractString(contents, commentCountRe))
	if err == nil {
		profile.CommentCount = commentCount
	}
	profile.Author = extractString(contents, authorRe)
	profile.DatePublished = extractString(contents, datePublishedRe)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp)  string{
	match := re.FindSubmatch(contents)
	if len(match) >=2 {
		return string(match[1])
	}else {
		return ""
	}
}