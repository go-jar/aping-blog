package resource

import (
	"blog/resource/log"
	"github.com/go-jar/mysql"

	"blog/conf"
)

var MysqlClientPool *mysql.Pool

func InitMysql() {
	config := &mysql.PoolConfig{NewClientFunc: NewMysqlClient}
	config.MaxConns = conf.MysqlConf.PoolSize
	config.MaxIdleTime = conf.MysqlConf.PoolClientMaxIdleTime

	MysqlClientPool = mysql.NewPool(config)
}

func NewMysqlClient() (*mysql.Client, error) {
	config := mysql.NewConfig(conf.MysqlConf.User, conf.MysqlConf.Pass, conf.MysqlConf.Host, conf.MysqlConf.Port, conf.MysqlConf.Name)
	config.LogLevel = conf.MysqlConf.LogLevel
	config.ReadTimeout = conf.MysqlConf.RWTimeout
	config.WriteTimeout = conf.MysqlConf.RWTimeout

	return mysql.NewClient(config, log.TraceLogger)
}
