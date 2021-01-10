package conf

type pprofConfJson struct {
	OnlineEnableHost string `json:"online_enable_host"`
	Port             string `json:"port"`
}

var PprofConf struct {
	Enable bool
	Port   string
}

func initPprofConf() {
	if serverConf.Pprof.OnlineEnableHost == CommonConf.Hostname {
		PprofConf.Enable = true
	} else {
		PprofConf.Enable = false
	}

	PprofConf.Port = serverConf.Pprof.Port
}
