package remark

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/errno"
)

func (rc *RemarkController) DeleteAction(context *RemarkContext) {
	if err := rc.VerifyToken(context.ApiContext); err != nil {
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return
	}

	id, err := rc.parseDeleteActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	if _, err := context.remarkSvc.DeleteById(id); err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"RequestId": context.TraceId,
	}
}

func (rc *RemarkController) parseDeleteActionParams(context *RemarkContext) (int64, *goerror.Error) {
	var id int64
	qs := query.NewQuerySet()
	qs.Int64Var(&id, "RemarkId", true, errno.ECommonInvalidArg, "invalid RemarkId", query.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("RemarkController.parseDeleteActionParams"), []byte(err.Error()))
		return -1, err
	}

	return id, nil
}
