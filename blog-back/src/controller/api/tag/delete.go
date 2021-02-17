package tag

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/errno"
)

func (tc *TagController) DeleteAction(context *TagContext) {
	id, err := tc.parseDeleteActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	if _, err := context.tagSvc.DeleteById(id); err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"RequestId": context.TraceId,
	}
}

func (tc *TagController) parseDeleteActionParams(context *TagContext) (int64, *goerror.Error) {
	var id int64
	qs := query.NewQuerySet()
	qs.Int64Var(&id, "Id", true, errno.ECommonInvalidArg, "invalid Id", query.CheckInt64IsPositive)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("TagController.parseCreateActionParams"), []byte(err.Error()))
		return -1, err
	}

	return id, nil
}
