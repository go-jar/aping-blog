package remark

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
		EntityName:     "remark",
		IdFieldName:    "Id",
	}
}

func (s *Svc) Insert(entities ...*entity.RemarkEntity) ([]int64, error) {
	var err error
	is := make([]interface{}, len(entities))
	for i, item := range entities {
		is[i] = item
	}

	ids, err := s.SqlRedis.Insert(s.EntityName, s.EntityName, s.IdFieldName, s.RedisKeyPrefix, CacheExpireSeconds, is...)
	if err != nil {
		s.ErrorLog([]byte("RemarkSvc.Insert"), []byte(err.Error()))
	}

	return ids, err
}

func (s *Svc) DeleteById(id int64) (bool, error) {
	find, err := s.SqlRedis.DeleteById(s.EntityName, s.EntityName, s.RedisKeyPrefix, id)
	if err != nil {
		s.ErrorLog([]byte("RemarkSvc.DeleteById"), []byte(err.Error()))
	}

	return find, err
}

func (s *Svc) UpdateById(id int64, newEntity *entity.RemarkEntity, updateFields map[string]bool) (bool, error) {
	setItems, err := s.SqlRedis.UpdateById(s.EntityName, s.EntityName, s.RedisKeyPrefix, id, newEntity, updateFields, CacheExpireSeconds)
	if err != nil {
		s.ErrorLog([]byte("RemarkSvc.UpdateById"), []byte(err.Error()))
	}

	return setItems != nil, err
}

func (s *Svc) GetRemarks(qp *mysql.QueryParams) ([]*entity.RemarkEntity, int64, error) {
	var entities []*entity.RemarkEntity

	err := s.SqlOrm.SimpleQueryAnd(s.EntityName, qp, entity.RemarkEntityType, &entities)
	if err != nil {
		s.ErrorLog([]byte("RemarkSvc.GetRemarks"), []byte(err.Error()))
		return nil, 0, err
	}

	total, err := s.SqlOrm.SimpleTotalAnd(s.EntityName, qp)
	if err != nil {
		s.ErrorLog([]byte("RemarkSvc.GetRemarks"), []byte(err.Error()))
		return nil, 0, err
	}

	return entities, total, nil
}

func (s *Svc) TotalGroupByArticleId() (map[int64]int64, error) {
	qb := new(mysql.QueryBuilder)
	qb.Select(s.EntityName, "article_id, count(*)").
		GroupBy("article_id")

	rows, err := s.SqlOrm.Dao().Client.Query(qb.Query(), qb.Args()...)
	defer s.SqlOrm.PutBackClient()
	if err != nil {
		s.ErrorLog([]byte("RemarkSvc.TotalGroupByArticleId"), []byte(err.Error()))
		return nil, err
	}

	result := map[int64]int64{}
	for rows.Next() {
		var articleId, remarkCount int64
		err = rows.Scan(&articleId, &remarkCount)
		if err != nil {
			s.ErrorLog([]byte("RemarkSvc.TotalGroupByArticleId"), []byte(err.Error()))
			return nil, err
		} else {
			result[articleId] = remarkCount
		}
	}

	return result, nil
}
