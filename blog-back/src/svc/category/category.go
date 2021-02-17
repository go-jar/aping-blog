package category

import (
	"blog/conf"
	"blog/entity"
	"blog/resource"
	"blog/svc"
	"github.com/go-jar/mysql"
	"github.com/go-jar/redis"
	"github.com/go-jar/sqlredis"
)

const (
	CacheExpireSeconds = 60 * 60 * 24
	TotalRowsCacheExpireSeconds = 60 * 10
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
		EntityName:     "category",
		IdFieldName:    "Id",
	}
}

func (s *Svc) Insert(entities ...*entity.CategoryEntity) ([]int64, error) {
	var err error
	is := make([]interface{}, len(entities))
	for i, entity := range entities {
		is[i] = entity
	}

	ids, err := s.SqlRedis.Insert(s.EntityName, s.EntityName, s.IdFieldName, s.RedisKeyPrefix, CacheExpireSeconds, is...)
	if err != nil {
		s.ErrorLog([]byte("Category.Insert"), []byte(err.Error()))
	}

	return ids, err
}

func (s *Svc) DeleteById(id int64) (bool, error) {
	find, err := s.SqlRedis.DeleteById(s.EntityName, s.EntityName, s.RedisKeyPrefix, id)
	if err != nil {
		s.ErrorLog([]byte("CategorySvc.DeleteById"), []byte(err.Error()))
	}

	return find, err
}

func (s *Svc) UpdateById(id int64, newEntity *entity.CategoryEntity, updateFields map[string]bool) (bool, error) {
	setItems, err := s.SqlRedis.UpdateById(s.EntityName, s.EntityName, s.RedisKeyPrefix, id, newEntity, updateFields, CacheExpireSeconds)
	if err != nil {
		s.ErrorLog([]byte("CategorySvc.UpdateById"), []byte(err.Error()))
	}

	return setItems != nil, err
}

func (s *Svc) GetById(id int64) (*entity.CategoryEntity, error) {
	user := new(entity.CategoryEntity)

	find, err := s.SqlRedis.GetById(s.EntityName, s.EntityName, s.RedisKeyPrefix, id, CacheExpireSeconds, user)
	if err != nil {
		s.ErrorLog([]byte("User.GetById"), []byte(err.Error()))
	}

	if !find {
		return nil, nil
	}

	return user, nil
}

func (s *Svc) SimpleQueryAnd(qp *mysql.QueryParams) ([]*entity.CategoryEntity, error) {
	var entities []*entity.CategoryEntity

	err := s.SqlOrm.SimpleQueryAnd(s.EntityName, qp, entity.CategoryEntityType, &entities)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (s *Svc) SimpleTotalAnd(qp *mysql.QueryParams) (int64, error) {
	return s.SqlOrm.SimpleTotalAnd(s.EntityName, qp)
}
