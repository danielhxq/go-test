package main

import (
	_ "awesomeProject3/seckill/secproxy/router"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func initConfig() {
	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")

	secKillConf.redisAddr = redisAddr
	secKillConf.etcdAddr = etcdAddr
	logs.Debug("dfdfd")
}

func main() {
	initConfig()
	beego.Run()
}
