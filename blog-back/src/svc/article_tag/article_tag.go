package article_tag

import (
	"time"

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
		EntityName:     "article_tag",
		IdFieldName:    "Id",
	}
}

func (s *Svc) Insert(articleId int64, tagIds ...int64) ([]int64, error) {
	is := make([]interface{}, len(tagIds))
	for i, tagId := range tagIds {
		is[i] = &entity.ArticleTagEntity{
			ArticleId:   articleId,
			TagId:       tagId,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}
	}

	ids, err := s.SqlRedis.Insert(s.EntityName, s.EntityName, s.IdFieldName, s.RedisKeyPrefix, CacheExpireSeconds, is...)
	if err != nil {
		s.ErrorLog([]byte("ArticleTagSvc.Insert"), []byte(err.Error()))
	}

	return ids, err
}

func (s *Svc) DeleteByArticleId(articleId int64) error {
	qb := new(mysql.QueryBuilder)
	qb.Delete(s.EntityName).
		WhereAnd(mysql.NewQueryItem("article_id", mysql.CondEqual, articleId))

	_, err := s.SqlOrm.Dao().Client.Exec(qb.Query(), qb.Args()...)
	defer s.SqlOrm.PutBackClient()
	if err != nil {
		return err
	}

	return nil
}

func (s *Svc) DeleteByTagIds(tagIds []int64) error {
	qb := new(mysql.QueryBuilder)
	qb.Delete(s.EntityName).
		WhereAnd(mysql.NewQueryItem("tag_id", mysql.CondIn, tagIds))

	_, err := s.SqlOrm.Dao().Client.Exec(qb.Query(), qb.Args()...)
	defer s.SqlOrm.PutBackClient()
	if err != nil {
		return err
	}

	return nil
}

func (s *Svc) UpdateArticleTags(articleId int64, newTagIds []int64) error {
	oldTagIds, err := s.GetTagIdsByArticleId(articleId)
	if err != nil {
		return err
	}

	newTagSet := s.listToMap(newTagIds)
	oldTagSet := s.listToMap(oldTagIds)

	var addTagIds []int64
	for _, newTagId := range newTagIds {
		if _, ok := oldTagSet[newTagId]; !ok {
			addTagIds = append(addTagIds, newTagId)
		}
	}

	var delTagIds []int64
	for _, oldTagId := range oldTagIds {
		if _, ok := newTagSet[oldTagId]; !ok {
			delTagIds = append(delTagIds, oldTagId)
		}
	}

	if len(addTagIds) > 0 {
		if _, err := s.Insert(articleId, addTagIds...); err != nil {
			return err
		}
	}

	if len(delTagIds) > 0 {
		if err := s.DeleteByTagIds(delTagIds); err != nil {
			return err
		}
	}

	return nil
}

func (s *Svc) listToMap(tagIds []int64) map[int64]bool {
	set := make(map[int64]bool)
	for _, oldTagId := range tagIds {
		set[oldTagId] = true
	}
	return set
}

func (s *Svc) SimpleQueryAnd(qp *mysql.QueryParams) ([]*entity.ArticleTagEntity, error) {
	var entities []*entity.ArticleTagEntity

	err := s.SqlOrm.SimpleQueryAnd(s.EntityName, qp, entity.ArticleTagEntityType, &entities)
	if err != nil {
		s.ErrorLog([]byte("ArticleTagSvc.GetArticles"), []byte(err.Error()))
		return nil, err
	}

	return entities, nil
}

func (s *Svc) GetTagIdsByArticleId(articleId int64) ([]int64, error) {
	qp := &mysql.QueryParams{
		ParamsStructPtr: &entity.ArticleTagEntity{
			ArticleId: articleId,
		},
		Required:   map[string]bool{"article_id": true},
		Conditions: map[string]string{"article_id": mysql.CondEqual},
	}
	articleTagEntities, err := s.SimpleQueryAnd(qp)
	if err != nil {
		return nil, err
	}

	var tagIds []int64
	for _, articleTagEntity := range articleTagEntities {
		tagIds = append(tagIds, articleTagEntity.TagId)
	}

	return tagIds, nil
}

func (s *Svc) GetByArticleIds(articleIds []int64) (map[int64][]int64, error) {
	qb := new(mysql.QueryBuilder)
	qb.Select(s.EntityName, "article_id, tag_id").
		WhereAnd(mysql.NewQueryItem("article_id", mysql.CondIn, articleIds))

	rows, err := s.SqlOrm.Dao().Client.Query(qb.Query(), qb.Args()...)
	defer s.SqlOrm.PutBackClient()
	if err != nil {
		return nil, err
	}

	result := map[int64][]int64{}
	for rows.Next() {
		var articleId, tagId int64
		err = rows.Scan(&articleId, &tagId)
		if err != nil {
			s.ErrorLog([]byte("ArticleTagSvc.GetByArticleIds"), []byte(err.Error()))
			return nil, err
		} else {
			result[articleId] = append(result[articleId], tagId)
		}
	}

	return result, nil
}

func (s *Svc) GetArticleIdsByTagId(tagId int64) ([]int64, error) {
	qp := &mysql.QueryParams{
		ParamsStructPtr: &entity.ArticleTagEntity{TagId: tagId},
		Required:   map[string]bool{"tag_id": true},
		Conditions: map[string]string{"tag_id": mysql.CondEqual},
	}
	articleTagEntities, err := s.SimpleQueryAnd(qp)
	if err != nil {
		return nil, err
	}

	var articleIds []int64
	for _, articleTagEntity := range articleTagEntities {
		articleIds = append(articleIds, articleTagEntity.ArticleId)
	}

	return articleIds, nil
}

func (s *Svc) TotalGroupByTagId() (map[int64]int64, error) {
	qb := new(mysql.QueryBuilder)
	qb.Select(s.EntityName, "tag_id, count(*)").
		GroupBy("tag_id")

	rows, err := s.SqlOrm.Dao().Client.Query(qb.Query(), qb.Args()...)
	defer s.SqlOrm.PutBackClient()
	if err != nil {
		s.ErrorLog([]byte("ArticleTagSvc.TotalGroupByTagId"), []byte(err.Error()))
		return nil, err
	}

	result := map[int64]int64{}
	for rows.Next() {
		var tagId, articleCount int64
		err = rows.Scan(&tagId, &articleCount)
		if err != nil {
			s.ErrorLog([]byte("ArticleTagSvc.TotalGroupByTagId"), []byte(err.Error()))
			return nil, err
		} else {
			result[tagId] = articleCount
		}
	}

	return result, nil
}
