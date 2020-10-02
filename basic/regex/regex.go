package main

import (
	"log"
	"regexp"
)

const text = "MY email is dd@gmail.com"

func main() {
	re, err := regexp.Compile("dd@gmail.com")
	if err != nil {
		panic(err)
	}
	match := re.FindString(text)
	log.Println(match)
}
