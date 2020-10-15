package main

var (
	secKillConf = &SecKillConf{}
)

type SecKillConf struct {
	redisAddr string
	etcdAddr  string
}
