package category

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/entity"
	"blog/errno"
)

func (cc *CategoryController) UpdateAction(context *CategoryContext) {
	categoryEntity, err := cc.parseUpdateActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	if _, err := context.categorySvc.UpdateById(categoryEntity.Id, categoryEntity, map[string]bool{"category_name": true}); err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{} {
		"RequestId": context.TraceId,
	}
}

func (cc *CategoryController) parseUpdateActionParams(context *CategoryContext) (*entity.CategoryEntity, *goerror.Error) {
	categoryEntity := &entity.CategoryEntity{}

	qs := query.NewQuerySet()
	qs.Int64Var(&categoryEntity.Id, "Id", true, errno.ECommonInvalidArg, "invalid Id", cc.CheckInt64GreaterEqual0)
	qs.StringVar(&categoryEntity.CategoryName, "CategoryName", true, errno.ECommonInvalidArg, "invalid CategoryName", query.CheckStringNotEmpty)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("CategoryController.parseCreateActionParams"), []byte(err.Error()))
		return nil, err
	}

	return categoryEntity, nil
}
