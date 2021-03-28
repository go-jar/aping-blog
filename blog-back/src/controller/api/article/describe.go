package article

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"
	"github.com/go-jar/mysql"

	"blog/entity"
	"blog/errno"
	"blog/svc/article_tag"
	"blog/svc/category"
	"blog/svc/remark"
	"blog/svc/tag"
)

func (ac *ArticleController) DescribeAction(context *ArticleContext) {
	qp, articleId, articleIds, findByArticleIds, err := ac.parseDescribeActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	var entities []*entity.ArticleEntity
	var total int64
	var e error

	if findByArticleIds {
		entities, total, e = context.articleSvc.GetArticlesByIdsLimit(articleIds, qp.Offset, qp.Cnt)
		if e != nil {
			context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
			return
		}
	} else {
		entities, total, e = context.articleSvc.GetArticles(qp)
		if e != nil {
			context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
			return
		}

		if articleId > -1 && len(entities) > 0 {
			articleEntity := entity.ArticleEntity{ReadCount: entities[0].ReadCount + 1}
			if _, err := context.articleSvc.UpdateById(articleId, &articleEntity, map[string]bool{"read_count": true}); err != nil {
				context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
				return
			}
		}
	}

	articleSet, e := ac.renderDescribeArticles(context.articleTagSvc, entities, context.TraceId)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.EInternalError, e.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"ArticleSet": articleSet,
		"Total":      total,
		"RequestId":  context.TraceId,
	}
}

func (ac *ArticleController) parseDescribeActionParams(context *ArticleContext) (*mysql.QueryParams, int64, []int64, bool, *goerror.Error) {
	qp := &mysql.QueryParams{OrderBy: "id desc"}
	var articleId, categoryId, tagId int64
	articleId = -1
	categoryId = -1
	tagId = -1
	keyword := ""
	findByArticleIds := false
	var articleIds []int64

	qs := query.NewQuerySet()
	qs.Int64Var(&qp.Offset, "Offset", false, errno.ECommonInvalidArg, "invalid Offset", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&qp.Cnt, "Limit", false, errno.ECommonInvalidArg, "invalid Cnt", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&articleId, "ArticleId", false, errno.ECommonInvalidArg, "invalid Id", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&categoryId, "CategoryId", false, errno.ECommonInvalidArg, "invalid CategoryId", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&tagId, "TagId", false, errno.ECommonInvalidArg, "invalid TagId", query.CheckInt64GreaterEqual0)
	qs.StringVar(&keyword, "Keyword", false, errno.ECommonInvalidArg, "invalid Keyword", query.CheckStringNotEmpty)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("ArticleController.parseDescribeActionParams"), []byte(err.Error()))
		return nil, articleId, nil, false, err
	}

	if articleId > -1 {
		qp.ParamsStructPtr = &entity.ArticleEntity{Id: articleId}
		qp.Required = map[string]bool{"id": true}
		qp.Conditions = map[string]string{"id": mysql.CondEqual}
	} else if categoryId > -1 {
		qp.ParamsStructPtr = &entity.ArticleEntity{CategoryId: categoryId}
		qp.Required = map[string]bool{"category_id": true}
		qp.Conditions = map[string]string{"category_id": mysql.CondEqual}
	} else if tagId > -1 {
		findByArticleIds = true
		var err error
		articleIds, err = context.articleTagSvc.GetArticleIdsByTagId(tagId)
		if err != nil {
			return nil, articleId, nil, findByArticleIds, goerror.New(errno.ESysMysqlError, "failed to get article ids by tag id")
		}
	} else if keyword != "" {
		qp.ParamsStructPtr = &entity.ArticleEntity{Title: "%" + keyword + "%",}
		qp.Required = map[string]bool{"title": true}
		qp.Conditions = map[string]string{"title": mysql.CondLike}
	}

	return qp, articleId, articleIds, findByArticleIds, nil
}

func (ac *ArticleController) renderDescribeArticles(articleTagSvc *article_tag.Svc, articles []*entity.ArticleEntity, traceId []byte) ([]*entity.ArticleResultEntity, error) {
	tagSvc := tag.NewSvc(traceId)
	tagMap, err := tagSvc.GetAll()
	if err != nil {
		return nil, err
	}

	categorySvc := category.NewSvc(traceId)
	categoryMap, err := categorySvc.GetAll()
	if err != nil {
		return nil, err
	}

	remarkSvc := remark.NewSvc(traceId)
	remarkMap, err := remarkSvc.TotalGroupByArticleId()
	if err != nil {
		return nil, err
	}

	var articleIds []int64
	for _, article := range articles {
		articleIds = append(articleIds, article.Id)
	}

	articleTagIdsMap, err := articleTagSvc.GetByArticleIds(articleIds)
	if err != nil {
		return nil, err
	}

	var articleSet []*entity.ArticleResultEntity
	for _, article := range articles {
		var tags []*entity.TagEntity
		if _, ok := articleTagIdsMap[article.Id]; ok {
			for _, tagId := range articleTagIdsMap[article.Id] {
				tags = append(tags, tagMap[tagId])
			}
		}

		articleSet = append(articleSet, &entity.ArticleResultEntity{
			Article:  article,
			Category: categoryMap[article.CategoryId],
			TagSet:   tags,
			RemarkCount: remarkMap[article.Id],
		})
	}

	return articleSet, nil
}
