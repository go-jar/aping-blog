package category

import (
	"github.com/go-jar/mysql"
	"github.com/go-jar/redis"
	"github.com/go-jar/sqlredis"

	"blog/conf"
	"blog/entity"
	"blog/resource"
	"blog/svc"
)

const (
	CacheExpireSeconds          = 60 * 60 * 24
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
	for i, item := range entities {
		is[i] = item
	}

	ids, err := s.SqlRedis.Insert(s.EntityName, s.EntityName, s.IdFieldName, s.RedisKeyPrefix, CacheExpireSeconds, is...)
	if err != nil {
		s.ErrorLog([]byte("CategorySvc.Insert"), []byte(err.Error()))
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

func (s *Svc) GetAll() (map[int64]*entity.CategoryEntity, error) {
	var categories []*entity.CategoryEntity
	qp := &mysql.QueryParams{Offset: 0, Cnt: 100000}
	err := s.SqlOrm.SimpleQueryAnd(s.EntityName, qp, entity.CategoryEntityType, &categories)
	if err != nil {
		s.ErrorLog([]byte("CategorySvc.GetAll"), []byte(err.Error()))
		return nil, err
	}

	result := map[int64]*entity.CategoryEntity{}
	for _, category := range categories {
		result[category.Id] = category
	}

	return result, nil
}

func (s *Svc) SimpleQueryAnd(qp *mysql.QueryParams) ([]*entity.CategoryEntity, error) {
	var entities []*entity.CategoryEntity

	err := s.SqlOrm.SimpleQueryAnd(s.EntityName, qp, entity.CategoryEntityType, &entities)
	if err != nil {
		s.ErrorLog([]byte("CategorySvc.SimpleQueryAnd"), []byte(err.Error()))
		return nil, err
	}

	return entities, nil
}

func (s *Svc) SimpleTotalAnd(qp *mysql.QueryParams) (int64, error) {
	total, err := s.SqlOrm.SimpleTotalAnd(s.EntityName, qp)
	if err != nil {
		s.ErrorLog([]byte("CategorySvc.SimpleTotalAnd"), []byte(err.Error()))
		return -1, err
	}

	return total, nil
}
