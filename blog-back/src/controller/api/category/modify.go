package category

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/entity"
	"blog/errno"
)

func (cc *CategoryController) ModifyAction(context *CategoryContext) {
	if err := cc.VerifyToken(context.ApiContext); err != nil {
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return
	}

	categoryEntity, err := cc.parseModifyActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	required := map[string]bool{
		"category_name":  true,
		"category_index": true,
	}
	if _, err := context.categorySvc.UpdateById(categoryEntity.Id, categoryEntity, required); err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"RequestId": context.TraceId,
	}
}

func (cc *CategoryController) parseModifyActionParams(context *CategoryContext) (*entity.CategoryEntity, *goerror.Error) {
	categoryEntity := &entity.CategoryEntity{}

	qs := query.NewQuerySet()
	qs.Int64Var(&categoryEntity.Id, "CategoryId", true, errno.ECommonInvalidArg, "invalid CategoryId", query.CheckInt64GreaterEqual0)
	qs.StringVar(&categoryEntity.CategoryName, "CategoryName", true, errno.ECommonInvalidArg, "invalid CategoryName", query.CheckStringNotEmpty)
	qs.Int64Var(&categoryEntity.CategoryIndex, "CategoryIndex", false, errno.ECommonInvalidArg, "invalid CategoryIndex", query.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("CategoryController.parseModifyActionParams"), []byte(err.Error()))
		return nil, err
	}

	return categoryEntity, nil
}
