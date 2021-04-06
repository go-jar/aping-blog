package remark

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"
	"github.com/go-jar/mysql"

	"blog/entity"
	"blog/errno"
	"blog/svc/remark"
)

func (rc *RemarkController) DescribeAction(context *RemarkContext) {
	qp, err := rc.parseDescribeActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	entities, total, e := context.remarkSvc.GetRemarks(qp)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
		return
	}

	remarkSet, e := rc.renderDescribeRemarks(context.remarkSvc, entities, context.TraceId)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"Total":     total,
		"RemarkSet": remarkSet,
		"RequestId": context.TraceId,
	}
}

func (rc *RemarkController) parseDescribeActionParams(context *RemarkContext) (*mysql.QueryParams, *goerror.Error) {
	qp := &mysql.QueryParams{OrderBy: "id desc"}
	var articleId int64
	articleId = -1

	qs := query.NewQuerySet()
	qs.Int64Var(&articleId, "ArticleId", false, errno.ECommonInvalidArg, "invalid ArticleId", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&qp.Offset, "Offset", false, errno.ECommonInvalidArg, "invalid Offset", query.CheckInt64GreaterEqual0)
	qs.Int64Var(&qp.Cnt, "Limit", false, errno.ECommonInvalidArg, "invalid Cnt", query.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("RemarkController.parseDescribeActionParams"), []byte(err.Error()))
		return nil, err
	}

	if articleId > -1 {
		qp.ParamsStructPtr = &entity.RemarkEntity{ArticleId: articleId}
		qp.Required = map[string]bool{"article_id": true}
		qp.Conditions = map[string]string{"article_id": mysql.CondEqual}
	}

	return qp, nil
}

func (rc *RemarkController) renderDescribeRemarks(remarkSvc *remark.Svc, remarks []*entity.RemarkEntity, traceId []byte) ([]*entity.RemarkResultEntity, error) {
	idRemarkMap := map[int64]*entity.RemarkEntity{}
	idReplyIdsMap := map[int64][]int64{}

	for _, remarkEntity := range remarks {
		idRemarkMap[remarkEntity.Id] = remarkEntity

		if remarkEntity.InitRemarkId != -1 {
			idReplyIdsMap[remarkEntity.InitRemarkId] = append([]int64{remarkEntity.Id}, idReplyIdsMap[remarkEntity.InitRemarkId]...)
		}
	}

	var remarkSet []*entity.RemarkResultEntity
	for _, remarkEntity := range remarks {
		if remarkEntity.InitRemarkId == 0 {
			var replies []*entity.RemarkEntity

			for _, remarkId := range idReplyIdsMap[remarkEntity.Id] {
				replies = append(replies, idRemarkMap[remarkId])
			}

			remarkSet = append(remarkSet, &entity.RemarkResultEntity {
				Remark:   remarkEntity,
				ReplySet: replies,
			})
		}
	}

	return remarkSet, nil
}
