package conf

import (
	"errors"

	"github.com/goinbox/gomisc"
)

var PrjHome string

func Init(prjHome string) error {
	if !gomisc.DirExist(prjHome) {
		return errors.New("prjHome not exists")
	}

	PrjHome = prjHome
	err := initServerConfJson()
	if err != nil {
		return errors.New("init serverConfJson error: " + err.Error())
	}

	initCommonConf()
	initLogConf()
	initPprofConf()
	initHttpConf()
	initMysqlConf()
	initRedisConf()
	initLoginConf()

	return nil
}
