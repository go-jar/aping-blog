package conf

type logConfJson struct {
	Level          int `json:"level"`
	AsyncQueueSize int `json:"async_queue_size"`
	BufSize        int `json:"buf_size"`
}

var LogConf struct {
	RootPath       string
	Level          int
	AsyncQueueSize int
	BufSize        int
}

func initLogConf() {
	LogConf.RootPath = PrjHome + "/logs"
	LogConf.Level = serverConf.Log.Level
	LogConf.AsyncQueueSize = serverConf.Log.AsyncQueueSize
	LogConf.BufSize = serverConf.Log.BufSize
}
