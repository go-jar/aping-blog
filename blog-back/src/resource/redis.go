package resource

import (
	"blog/resource/log"
	"github.com/go-jar/redis"

	"blog/conf"
)

var RedisClientPoolList []*redis.Pool

func InitRedis() {
	for _, rconf := range conf.RedisConfList {
		config := &redis.PoolConfig{NewClientFunc: NewRedisClientFunc(rconf)}
		config.MaxConns = rconf.PoolSize
		config.MaxIdleTime = rconf.PoolClientMaxIdleTime
		config.KeepAliveInterval = rconf.PoolKeepAliveInterval

		RedisClientPoolList = append(RedisClientPoolList, redis.NewPool(config))
	}
}

func NewRedisClientFunc(rconf *conf.RedisConf) func() (*redis.Client, error) {
	return func() (*redis.Client, error) {
		config := redis.NewConfig(rconf.Host, rconf.Port, rconf.Pass)
		config.LogLevel = rconf.LogLevel
		config.ConnectTimeout = rconf.ConnectTimeout
		config.ReadTimeout = rconf.RWTimeout
		config.WriteTimeout = rconf.RWTimeout

		return redis.NewClient(config, log.TraceLogger), nil
	}
}
