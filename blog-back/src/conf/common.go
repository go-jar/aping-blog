package conf

import (
	"os"
	"os/user"
)

var CommonConf struct {
	Hostname string
	Username string

	PrjName string
	IsDev   bool
	Idc     string

	TmpRoot    string
	ApiPidFile string
}

func initCommonConf() {
	CommonConf.Hostname, _ = os.Hostname()
	curUser, _ := user.Current()
	CommonConf.Username = curUser.Username

	CommonConf.PrjName = serverConf.PrjName
	CommonConf.IsDev = serverConf.IsDev
	CommonConf.Idc = serverConf.Idc

	CommonConf.TmpRoot = PrjHome + "/tmp"
	CommonConf.ApiPidFile = CommonConf.TmpRoot + "/api.pid"
}
