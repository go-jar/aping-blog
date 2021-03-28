package conf

import "blog/utils"

var LoginSecretKey string
var LoginExpireSeconds int64
var AdminName string

func initLoginConf() {
	LoginSecretKey = utils.GenSecretKey(12)
	LoginExpireSeconds = 60 * 60 * 24 * 3
	AdminName = "admin"
}
