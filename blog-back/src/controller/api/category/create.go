package category

import (
	"time"

	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/entity"
	"blog/errno"
)

func (cc *CategoryController) CreateAction(context *CategoryContext) {
	if err := cc.VerifyToken(context.ApiContext); err != nil {
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return
	}

	categoryEntity, e := cc.parseCreateActionParams(context)
	if e != nil {
		context.ApiData.Err = e
		return
	}

	ids, err := context.categorySvc.Insert(categoryEntity)
	if err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"CategoryId":        ids[0],
		"RequestId": context.TraceId,
	}
}

func (cc *CategoryController) parseCreateActionParams(context *CategoryContext) (*entity.CategoryEntity, *goerror.Error) {
	categoryEntity := &entity.CategoryEntity{
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	qs := query.NewQuerySet()
	qs.StringVar(&categoryEntity.CategoryName, "CategoryName", true, errno.ECommonInvalidArg, "invalid CategoryName", query.CheckStringNotEmpty)
	qs.Int64Var(&categoryEntity.CategoryIndex, "CategoryIndex", false, errno.ECommonInvalidArg, "invalid CategoryIndex", query.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("CategoryController.parseCreateActionParams"), []byte(err.Error()))
		return nil, err
	}

	return categoryEntity, nil
}
