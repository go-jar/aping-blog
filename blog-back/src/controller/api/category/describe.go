package category

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"
	"github.com/go-jar/mysql"

	"blog/entity"
	"blog/errno"
	"blog/svc/article"
)

func (cc *CategoryController) DescribeAction(context *CategoryContext) {
	qp, err := cc.parseDescribeActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	entities, e := context.categorySvc.SimpleQueryAnd(qp)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
		return
	}

	categorySet, e := cc.renderDescribeCategories(context, entities)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.EInternalError, e.Error())
		return
	}

	total, e := context.categorySvc.SimpleTotalAnd(qp)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"CategorySet": categorySet,
		"Total":       total,
		"RequestId":   context.TraceId,
	}
}

func (cc *CategoryController) parseDescribeActionParams(context *CategoryContext) (*mysql.QueryParams, *goerror.Error) {
	qp := &mysql.QueryParams{OrderBy: "category_index"}
	var id int64
	id = -1

	qs := query.NewQuerySet()
	qs.Int64Var(&qp.Offset, "Offset", false, errno.ECommonInvalidArg, "invalid Offset", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&qp.Cnt, "Limit", false, errno.ECommonInvalidArg, "invalid Cnt", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&id, "CategoryId", false, errno.ECommonInvalidArg, "invalid CategoryId", query.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("CategoryController.parseDescribeActionParams"), []byte(err.Error()))
		return nil, err
	}

	if id > -1 {
		qp.ParamsStructPtr = &entity.CategoryEntity{Id: id}
		qp.Required = map[string]bool{"id": true}
		qp.Conditions = map[string]string{"id": mysql.CondEqual}
	}

	return qp, nil
}

func (cc *CategoryController) renderDescribeCategories(context *CategoryContext, categories []*entity.CategoryEntity) ([]*entity.CategoryResultEntity, error) {
	articleSvc := article.NewSvc(context.TraceId)
	m, err := articleSvc.TotalGroupByCategoryId()
	if err != nil {
		return nil, err
	}
	
	var categorySet []*entity.CategoryResultEntity
	for _, category := range categories {
		if _, ok := m[category.Id]; ok {
			categorySet = append(categorySet, &entity.CategoryResultEntity{
				Category:     category,
				ArticleCount: m[category.Id],
			})
		} else {
			categorySet = append(categorySet, &entity.CategoryResultEntity{
				Category:     category,
				ArticleCount: 0,
			})
		}
	}

	return categorySet, nil
}
