package category

import (
	"fmt"

	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"
	"github.com/go-jar/mysql"

	"blog/entity"
	"blog/errno"
)

func (cc *CategoryController) DescribeAction(context *CategoryContext) {
	qp, err := cc.parseDescribeActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	entities, e := context.categorySvc.SimpleQueryAnd(qp)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
		return
	}

	total, e := context.categorySvc.SimpleTotalAnd(qp)
	if e != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, e.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"CategorySet": entities,
		"Total": total,
		"RequestId": context.TraceId,
	}
}

func (cc *CategoryController) parseDescribeActionParams(context *CategoryContext) (*mysql.QueryParams, *goerror.Error) {
	qp := &mysql.QueryParams{OrderBy: "id desc"}
	var id int64
	id = -1

	qs := query.NewQuerySet()
	qs.Int64Var(&qp.Offset, "Offset", false, errno.ECommonInvalidArg, "invalid Offset", cc.CheckInt64GreaterEqual0)
	qs.Int64Var(&qp.Cnt, "Limit", false, errno.ECommonInvalidArg, "invalid Cnt", cc.CheckInt64GreaterEqual0)
	qs.Int64Var(&id, "Id", false, errno.ECommonInvalidArg, "invalid Id", cc.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("CategoryController.parseDescribeActionParams"), []byte(err.Error()))
		return nil, err
	}

	if id > -1 {
		qp.ParamsStructPtr = &entity.CategoryEntity{Id: id}
		qp.Required = map[string]bool{"id": true}
		qp.Conditions = map[string]string{"id": mysql.CondEqual}
	}
	fmt.Println(qp)
	return qp, nil
}

func (cc *CategoryController) CheckInt64GreaterEqual0(v int64) bool {
	if v >= 0 {
		return true
	}
	return false
}
