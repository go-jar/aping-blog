package category

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/errno"
)

func (cc *CategoryController) DeleteAction(context *CategoryContext) {
	id, err := cc.parseDeleteActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	if _, err := context.categorySvc.DeleteById(id); err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"RequestId": context.TraceId,
	}
}

func (cc *CategoryController) parseDeleteActionParams(context *CategoryContext) (int64, *goerror.Error) {
	var id int64
	qs := query.NewQuerySet()
	qs.Int64Var(&id, "Id", true, errno.ECommonInvalidArg, "invalid Id", query.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("CategoryController.parseCreateActionParams"), []byte(err.Error()))
		return -1, err
	}

	return id, nil
}
