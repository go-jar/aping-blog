package user

import (
	"fmt"
	"reflect"

	"github.com/go-jar/crypto"
	"github.com/go-jar/mysql"
	"github.com/go-jar/redis"
	"github.com/go-jar/sqlredis"

	"blog/conf"
	"blog/entity"
	"blog/resource"
	"blog/svc"
)

type Svc struct {
	*svc.BaseSvc
	*sqlredis.SqlRedis

	RedisKeyPrefix string
	EntityName     string
}

func NewSvc(traceId []byte) *Svc {
	return &Svc{
		BaseSvc: &svc.BaseSvc{
			TraceId: traceId,
		},
		SqlRedis: &sqlredis.SqlRedis{
			SqlOrm:   mysql.NewSimpleOrm(traceId, resource.MysqlClientPool, false),
			RedisOrm: redis.NewSimpleOrm(traceId, resource.RedisClientPoolList[0]),
		},

		RedisKeyPrefix: conf.CommonConf.PrjName,
		EntityName:     "user",
	}
}

func (s *Svc) GetByUsername(username string) (*entity.UserEntity, error) {
	var users []*entity.UserEntity
	qp := &mysql.QueryParams{
		ParamsStructPtr: &entity.UserEntity{
			Username: username,
		},
		Required:   map[string]bool{"username": true},
		Conditions: map[string]string{"username": mysql.CondEqual},
		OrderBy:    "",
	}

	err := s.SqlOrm.SimpleQueryAnd(s.EntityName, qp, reflect.TypeOf(entity.UserEntity{}), &users)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("there no user named %s", username)
	}

	return users[0], nil
}

func (s *Svc) VerifyPassword(passwordFromRequest, passwordFromDb string) bool {
	if crypto.Md5String([]byte(passwordFromRequest)) != passwordFromDb {
		return false
	}
	return true
}
