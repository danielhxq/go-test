package main

import (
	_ "github.com/Shopify/sarama"
	"github.com/astaxie/beego/config"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/hpcloud/tail"
	"log"
)

func main() {
	_, err := config.NewConfig("yaml", "test.yml")
	if err != nil {
		log.Fatal(err)
	}
}
