package parser

import (
	"awesomeProject3/crawler/engine"
	"awesomeProject3/model"
	"regexp"
)

const ageRe = `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`

func ParseProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(ageRe)
	all := re.FindAllSubmatch(contents, -1)

	profile := model.Profile{}

	len := len(all)

	profile.Marriage = string(all[len][1])

	result := engine.ParseResult{Items: []interface{}{profile}}
	return result
}
