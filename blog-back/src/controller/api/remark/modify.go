package remark

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/entity"
	"blog/errno"
)

func (rc *RemarkController) ModifyAction(context *RemarkContext) {
	if err := rc.VerifyToken(context.ApiContext); err != nil {
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return
	}

	remarkEntity, err := rc.parseModifyActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	required := map[string]bool {
		"content":  true,
	}
	if _, err := context.remarkSvc.UpdateById(remarkEntity.Id, remarkEntity, required); err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"RequestId": context.TraceId,
	}
}

func (rc *RemarkController) parseModifyActionParams(context *RemarkContext) (*entity.RemarkEntity, *goerror.Error) {
	remarkEntity := &entity.RemarkEntity{}

	qs := query.NewQuerySet()
	qs.Int64Var(&remarkEntity.Id, "RemarkId", true, errno.ECommonInvalidArg, "invalid RemarkId", query.CheckInt64GreaterEqual0)
	qs.StringVar(&remarkEntity.Content, "Content", true, errno.ECommonInvalidArg, "invalid Content", query.CheckStringNotEmpty)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("RemarkController.parseModifyActionParams"), []byte(err.Error()))
		return nil, err
	}

	return remarkEntity, nil
}
