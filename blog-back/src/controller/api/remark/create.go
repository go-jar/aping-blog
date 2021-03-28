package remark

import (
	"time"

	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/entity"
	"blog/errno"
)

func (rc *RemarkController) CreateAction(context *RemarkContext) {
	remarkEntity, e := rc.parseCreateActionParams(context)
	if e != nil {
		context.ApiData.Err = e
		return
	}

	ids, err := context.remarkSvc.Insert(remarkEntity)
	if err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"RemarkId":        ids[0],
		"RequestId": context.TraceId,
	}
}

func (rc *RemarkController) parseCreateActionParams(context *RemarkContext) (*entity.RemarkEntity, *goerror.Error) {
	remarkEntity := &entity.RemarkEntity{
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	qs := query.NewQuerySet()
	qs.Int64Var(&remarkEntity.ArticleId, "ArticleId", true, errno.ECommonInvalidArg, "invalid ArticleId", query.CheckInt64GreaterEqual0)
	qs.StringVar(&remarkEntity.Nickname, "Nickname", true, errno.ECommonInvalidArg, "invalid Nickname", query.CheckStringNotEmpty)
	qs.StringVar(&remarkEntity.Content, "Content", true, errno.ECommonInvalidArg, "invalid Content", query.CheckStringNotEmpty)
	qs.Int64Var(&remarkEntity.InitRemarkId, "InitRemarkId", false, errno.ECommonInvalidArg, "invalid InitRemarkId", query.CheckInt64GreaterEqual0)
	qs.StringVar(&remarkEntity.NicknameReplied, "NicknameReplied", false, errno.ECommonInvalidArg, "invalid NicknameReplied", query.CheckStringNotEmpty)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("RemarkController.parseCreateActionParams"), []byte(err.Error()))
		return nil, err
	}

	return remarkEntity, nil
}
