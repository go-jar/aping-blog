package article

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
		EntityName:     "article",
		IdFieldName:    "Id",
	}
}

func (s *Svc) Insert(entities ...*entity.ArticleEntity) ([]int64, error) {
	var err error
	is := make([]interface{}, len(entities))
	for i, item := range entities {
		is[i] = item
	}

	ids, err := s.SqlRedis.Insert(s.EntityName, s.EntityName, s.IdFieldName, s.RedisKeyPrefix, CacheExpireSeconds, is...)
	if err != nil {
		s.ErrorLog([]byte("ArticleSvc.Insert"), []byte(err.Error()))
	}

	return ids, err
}

func (s *Svc) DeleteById(id int64) (bool, error) {
	find, err := s.SqlRedis.DeleteById(s.EntityName, s.EntityName, s.RedisKeyPrefix, id)
	if err != nil {
		s.ErrorLog([]byte("ArticleSvc.DeleteById"), []byte(err.Error()))
	}

	return find, err
}

func (s *Svc) UpdateById(id int64, newEntity *entity.ArticleEntity, updateFields map[string]bool) (bool, error) {
	setItems, err := s.SqlRedis.UpdateById(s.EntityName, s.EntityName, s.RedisKeyPrefix, id, newEntity, updateFields, CacheExpireSeconds)
	if err != nil {
		s.ErrorLog([]byte("ArticleSvc.UpdateById"), []byte(err.Error()))
	}

	return setItems != nil, err
}

func (s *Svc) GetArticles(qp *mysql.QueryParams) ([]*entity.ArticleEntity, int64, error) {
	var entities []*entity.ArticleEntity

	err := s.SqlOrm.SimpleQueryAnd(s.EntityName, qp, entity.ArticleEntityType, &entities)
	if err != nil {
		s.ErrorLog([]byte("ArticleSvc.GetArticles"), []byte(err.Error()))
		return nil, 0, err
	}

	total, err := s.SqlOrm.SimpleTotalAnd(s.EntityName, qp)
	if err != nil {
		s.ErrorLog([]byte("ArticleSvc.GetArticles"), []byte(err.Error()))
		return nil, 0, err
	}

	return entities, total, nil
}

func (s *Svc) TotalGroupByCategoryId() (map[int64]int64, error) {
	qb := new(mysql.QueryBuilder)
	qb.Select(s.EntityName, "category_id, count(*)").
		GroupBy("category_id")

	rows, err := s.SqlOrm.Dao().Client.Query(qb.Query(), qb.Args()...)
	defer s.SqlOrm.PutBackClient()
	if err != nil {
		s.ErrorLog([]byte("ArticleSvc.TotalGroupByCategoryId"), []byte(err.Error()))
		return nil, err
	}

	result := map[int64]int64{}
	for rows.Next() {
		var categoryId, articleCount int64
		err = rows.Scan(&categoryId, &articleCount)
		if err != nil {
			s.ErrorLog([]byte("ArticleSvc.TotalGroupByCategoryId"), []byte(err.Error()))
			return nil, err
		} else {
			result[categoryId] = articleCount
		}
	}

	return result, nil
}

func (s *Svc) GetArticlesByIdsLimit(articleIds []int64, offset, limit int64) ([]*entity.ArticleEntity, int64, error) {
	var entities []*entity.ArticleEntity

	if articleIds == nil || len(articleIds) == 0 {
		return entities, 0, nil
	}

	if err := s.SqlOrm.ListByIdsLimit(s.EntityName, articleIds, "id desc", offset, limit, entity.ArticleEntityType, &entities); err != nil {
		s.ErrorLog([]byte("ArticleSvc.GetArticlesByIdsLimit"), []byte(err.Error()))
		return nil, -1, err
	}

	total := int64(len(articleIds))

	return entities, total, nil
}
