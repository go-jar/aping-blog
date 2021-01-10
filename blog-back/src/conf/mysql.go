package conf

import (
	"fmt"
	"time"

	"blog/utils"
)

type mysqlConfJson struct {
	Host                     string `json:"host"`
	User                     string `json:"user"`
	Pass                     string `json:"pass"`
	Port                     string `json:"port"`
	Name                     string `json:"name"`
	LogLevel                 int    `json:"log_level"`
	RWTimeoutSeconds         int    `json:"rw_timeout_seconds"`
	PoolSize                 int    `json:"pool_size"`
	PoolClientMaxIdleSeconds int    `json:"pool_client_max_idle_seconds"`
}

var MysqlConf struct {
	Host                  string
	User                  string
	Pass                  string
	Port                  string
	Name                  string
	LogLevel              int
	RWTimeout             time.Duration
	PoolSize              int
	PoolClientMaxIdleTime time.Duration
}

func initMysqlConf() {
	MysqlConf.Host = serverConf.Mysql.Host
	MysqlConf.User = serverConf.Mysql.User
	fmt.Println("initMysqlConf ", utils.DecryptString(serverConf.Mysql.Pass))
	MysqlConf.Pass = utils.DecryptString(serverConf.Mysql.Pass)
	MysqlConf.Port = serverConf.Mysql.Port
	MysqlConf.Name = serverConf.Mysql.Name
	MysqlConf.LogLevel = serverConf.Mysql.LogLevel
	MysqlConf.RWTimeout = time.Duration(serverConf.Mysql.RWTimeoutSeconds) * time.Second
	MysqlConf.PoolSize = serverConf.Mysql.PoolSize
	MysqlConf.PoolClientMaxIdleTime = time.Duration(serverConf.Mysql.PoolClientMaxIdleSeconds) * time.Second
}
