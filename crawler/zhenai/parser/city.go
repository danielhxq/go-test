package parser

import (
	"awesomeProject3/crawler/engine"
	"regexp"
)

const city = `<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(city)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range all {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: ParseProfile})
	}
	return result
}
