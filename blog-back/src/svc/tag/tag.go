package tag

import (
	"github.com/go-jar/mysql"
	"github.com/go-jar/redis"
	"github.com/go-jar/sqlredis"
	"reflect"

	"blog/conf"
	"blog/entity"
	"blog/resource"
	"blog/svc"
)

const (
	CacheExpireSeconds = 60 * 60 * 24
)

type Svc struct {
	*svc.BaseSvc
	*sqlredis.SqlRedis

	RedisKeyPrefix string
	EntityName     string
	IdFieldName    string
}

func NewSvc(traceId []byte) *Svc {
	return &Svc{
		BaseSvc: &svc.BaseSvc{
			TraceId: traceId,
		},
		SqlRedis: &sqlredis.SqlRedis{
			SqlOrm:   mysql.NewSimpleOrm(traceId, resource.MysqlClientPool, true),
			RedisOrm: redis.NewSimpleOrm(traceId, resource.RedisClientPoolList[0]),
		},

		RedisKeyPrefix: conf.CommonConf.PrjName,
		EntityName:     "tag",
		IdFieldName:    "Id",
	}
}

func (s *Svc) Insert(entities ...*entity.TagEntity) ([]int64, error) {
	var err error
	is := make([]interface{}, len(entities))
	for i, item := range entities {
		is[i] = item
	}

	ids, err := s.SqlRedis.Insert(s.EntityName, s.EntityName, s.IdFieldName, s.RedisKeyPrefix, CacheExpireSeconds, is...)
	if err != nil {
		s.ErrorLog([]byte("TagSvc.Insert"), []byte(err.Error()))
	}

	return ids, err
}

func (s *Svc) DeleteById(id int64) (bool, error) {
	find, err := s.SqlRedis.DeleteById(s.EntityName, s.EntityName, s.RedisKeyPrefix, id)
	if err != nil {
		s.ErrorLog([]byte("TagSvc.DeleteById"), []byte(err.Error()))
	}

	return find, err
}

func (s *Svc) UpdateById(id int64, newEntity *entity.TagEntity, updateFields map[string]bool) (bool, error) {
	setItems, err := s.SqlRedis.UpdateById(s.EntityName, s.EntityName, s.RedisKeyPrefix, id, newEntity, updateFields, CacheExpireSeconds)
	if err != nil {
		s.ErrorLog([]byte("TagSvc.UpdateById"), []byte(err.Error()))
	}

	return setItems != nil, err
}

func (s *Svc) GetByIds(ids []int64) ([]*entity.TagEntity, error) {
	var data []*entity.TagEntity
	tagEntityType := reflect.TypeOf(entity.TagEntity{})

	if err := s.SqlOrm.ListByIds(s.EntityName, ids, "", tagEntityType, &data); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (s *Svc) SimpleQueryAnd(qp *mysql.QueryParams) ([]*entity.TagEntity, error) {
	var entities []*entity.TagEntity

	err := s.SqlOrm.SimpleQueryAnd(s.EntityName, qp, entity.TagEntityType, &entities)
	if err != nil {
		s.ErrorLog([]byte("TagSvc.SimpleQueryAnd"), []byte(err.Error()))
		return nil, err
	}

	return entities, nil
}

func (s *Svc) SimpleTotalAnd(qp *mysql.QueryParams) (int64, error) {
	total, err := s.SqlOrm.SimpleTotalAnd(s.EntityName, qp)
	if err != nil {
		s.ErrorLog([]byte("TagSvc.SimpleTotalAnd"), []byte(err.Error()))
		return -1, err
	}

	return total, nil
}
