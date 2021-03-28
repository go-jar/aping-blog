package tag

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"
	"github.com/go-jar/mysql"

	"blog/entity"
	"blog/errno"
	"blog/svc/article_tag"
)

func (tc *TagController) DescribeAction(context *TagContext) {
	qp, err := tc.parseDescribeActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	entities, e := context.tagSvc.SimpleQueryAnd(qp)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
		return
	}

	tagSet, e := tc.renderDescribeTags(context, entities)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.EInternalError, e.Error())
		return
	}

	total, e := context.tagSvc.SimpleTotalAnd(qp)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"TagSet":    tagSet,
		"Total":     total,
		"RequestId": context.TraceId,
	}
}

func (tc *TagController) parseDescribeActionParams(context *TagContext) (*mysql.QueryParams, *goerror.Error) {
	qp := &mysql.QueryParams{OrderBy: "tag_index"}
	var id int64
	id = -1

	qs := query.NewQuerySet()
	qs.Int64Var(&qp.Offset, "Offset", false, errno.ECommonInvalidArg, "invalid Offset", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&qp.Cnt, "Limit", false, errno.ECommonInvalidArg, "invalid Cnt", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&id, "TagId", false, errno.ECommonInvalidArg, "invalid TagId", query.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("TagController.parseDescribeActionParams"), []byte(err.Error()))
		return nil, err
	}

	if id > -1 {
		qp.ParamsStructPtr = &entity.CategoryEntity{Id: id}
		qp.Required = map[string]bool{"id": true}
		qp.Conditions = map[string]string{"id": mysql.CondEqual}
	}

	return qp, nil
}

func (tc *TagController) renderDescribeTags(context *TagContext, tags []*entity.TagEntity) ([]*entity.TagResultEntity, error) {
	articleTagSvc := article_tag.NewSvc(context.TraceId)
	m, err := articleTagSvc.TotalGroupByTagId()
	if err != nil {
		return nil, err
	}

	var tagSet []*entity.TagResultEntity
	for _, tag := range tags {
		if _, ok := m[tag.Id]; ok {
			tagSet = append(tagSet, &entity.TagResultEntity{
				Tag:          tag,
				ArticleCount: m[tag.Id],
			})
		} else {
			tagSet = append(tagSet, &entity.TagResultEntity{
				Tag:          tag,
				ArticleCount: 0,
			})
		}
	}

	return tagSet, nil
}
