package conf

import "blog/utils"

var LoginSecretKey string
var LoginExpireSeconds int64

func initLoginConf() {
	LoginSecretKey = utils.GenSecretKey(12)
	LoginExpireSeconds = 60 * 60 * 24
}
