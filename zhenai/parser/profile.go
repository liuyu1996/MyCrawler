package parser

import (
	"crawler1/engine"
	"crawler1/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d])+岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<])+</div>`)
//var genderRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<])+</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d])+cm</div>`)
var wightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d])+kg</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<])+</div>`)
var nameRe = regexp.MustCompile(`<h1 class="nickName" data-v-5b109fc3>([^<])+</h1>`)
//var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<])+</div>`)
//var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<])+</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>籍贯:([^<])+</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^\(])+</div>`)

func ParseProfile(contents []byte)  engine.ParseResult{
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	hight, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Height = hight
	}
	wight, err := strconv.Atoi(extractString(contents, wightRe))
	if err != nil {
		profile.Weight = wight
	}
	profile.Marriage = extractString(contents, marriageRe)
	profile.Name = extractString(contents, nameRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Income = extractString(contents, incomeRe)

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