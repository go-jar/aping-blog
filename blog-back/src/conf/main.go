package conf

import (
	"github.com/goinbox/gomisc"
)

var serverConf serverConfJson

type serverConfJson struct {
	PrjName string `json:"prj_name"`
	IsDev   bool   `json:"is_dev"`
	Idc     string `json:"idc"`

	Log     logConfJson   `json:"log"`
	Pprof   pprofConfJson `json:"pprof"`
	ApiHttp httpConfJson  `json:"api_http"`

	Mysql mysqlConfJson    `json:"mysql"`
	Redis []*redisConfJson `json:"redis"`

	TcloudAccount tcloudAccountJson  `json:"tcloud_account"`
}

func initServerConfJson() error {
	confRoot := PrjHome + "/conf"
	err := gomisc.ParseJsonFile(confRoot+"/server/server_conf.json", &serverConf)
	if err != nil {
		return err
	}

	return nil
}
