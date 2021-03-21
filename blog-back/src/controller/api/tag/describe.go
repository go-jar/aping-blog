package tag

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"
	"github.com/go-jar/mysql"

	"blog/entity"
	"blog/errno"
)

func (tc *TagController) DescribeAction(context *TagContext) {
	qp, err := tc.parseDescribeActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	entities, e := context.tagSvc.SimpleQueryAnd(qp)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
		return
	}

	total, e := context.tagSvc.SimpleTotalAnd(qp)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"TagSet":    entities,
		"Total":     total,
		"RequestId": context.TraceId,
	}
}

func (tc *TagController) parseDescribeActionParams(context *TagContext) (*mysql.QueryParams, *goerror.Error) {
	qp := &mysql.QueryParams{OrderBy: "tag_index"}
	var id int64
	id = -1

	qs := query.NewQuerySet()
	qs.Int64Var(&qp.Offset, "Offset", false, errno.ECommonInvalidArg, "invalid Offset", tc.CheckInt64GreaterEqual0)
	qs.Int64Var(&qp.Cnt, "Limit", false, errno.ECommonInvalidArg, "invalid Cnt", tc.CheckInt64GreaterEqual0)
	qs.Int64Var(&id, "Id", false, errno.ECommonInvalidArg, "invalid Id", tc.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("TagController.parseDescribeActionParams"), []byte(err.Error()))
		return nil, err
	}

	if id > -1 {
		qp.ParamsStructPtr = &entity.CategoryEntity{Id: id}
		qp.Required = map[string]bool{"id": true}
		qp.Conditions = map[string]string{"id": mysql.CondEqual}
	}
	
	return qp, nil
}

func (tc *TagController) CheckInt64GreaterEqual0(v int64) bool {
	if v >= 0 {
		return true
	}
	return false
}
