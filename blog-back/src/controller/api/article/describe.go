package article

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"
	"github.com/go-jar/mysql"

	"blog/entity"
	"blog/errno"
	"blog/svc/article_tag"
	"blog/svc/category"
	"blog/svc/tag"
)

func (ac *ArticleController) DescribeAction(context *ArticleContext) {
	qp, ids, findByIds, err := ac.parseDescribeActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	var entities []*entity.ArticleEntity
	var total int64
	var e error

	if findByIds {
		entities, total, e = context.articleSvc.GetArticlesByIdsLimit(ids, qp.Offset, qp.Cnt)
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

func (ac *ArticleController) parseDescribeActionParams(context *ArticleContext) (*mysql.QueryParams, []int64, bool, *goerror.Error) {
	qp := &mysql.QueryParams{OrderBy: "id desc"}
	var articleId, categoryId, tagId int64
	articleId = -1
	categoryId = -1
	tagId = -1

	qs := query.NewQuerySet()
	qs.Int64Var(&qp.Offset, "Offset", false, errno.ECommonInvalidArg, "invalid Offset", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&qp.Cnt, "Limit", false, errno.ECommonInvalidArg, "invalid Cnt", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&articleId, "Id", false, errno.ECommonInvalidArg, "invalid Id", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&categoryId, "CategoryId", false, errno.ECommonInvalidArg, "invalid CategoryId", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&tagId, "TagId", false, errno.ECommonInvalidArg, "invalid TagId", query.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("ArticleController.parseDescribeActionParams"), []byte(err.Error()))
		return nil, nil, false, err
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
		articleIds, err := context.articleTagSvc.GetArticleIdsByTagId(tagId)
		if err != nil {
			return nil, nil, true, goerror.New(errno.ESysMysqlError, "failed to get article ids by tag id")
		}
		return qp, articleIds, true, nil
	}

	return qp, nil, false, nil
}

func (ac *ArticleController) renderDescribeArticles(articleTagSvc *article_tag.Svc, articles []*entity.ArticleEntity, traceId []byte) ([]*entity.ArticleResultEntity, error) {
	var articleSet []*entity.ArticleResultEntity

	for _, articleEntity := range articles {
		tagIds, err := articleTagSvc.GetTagIdsByArticleId(articleEntity.Id)
		if err != nil {
			return nil, err
		}

		var tagEntities []*entity.TagEntity
		if len(tagIds) > 0 {
			tagSvc := tag.NewSvc(traceId)
			tagEntities, err = tagSvc.GetByIds(tagIds)
			if err != nil {
				return nil, err
			}
		}

		categorySvc := category.NewSvc(traceId)
		categoryEntity, err := categorySvc.GetById(articleEntity.CategoryId)
		if err != nil {
			return nil, err
		}

		articleSet = append(articleSet, &entity.ArticleResultEntity{
			Article:  articleEntity,
			Category: categoryEntity,
			TagSet:   tagEntities,
		})
	}

	return articleSet, nil
}
