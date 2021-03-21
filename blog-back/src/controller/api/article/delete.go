package article

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/errno"
)

func (ac *ArticleController) DeleteAction(context *ArticleContext) {
	id, err := ac.parseDeleteActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	if _, err := context.articleSvc.DeleteById(id); err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	if err := context.articleTagSvc.DeleteByArticleId(id); err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"RequestId": context.TraceId,
	}
}

func (ac *ArticleController) parseDeleteActionParams(context *ArticleContext) (int64, *goerror.Error) {
	var id int64
	qs := query.NewQuerySet()
	qs.Int64Var(&id, "Id", true, errno.ECommonInvalidArg, "invalid Id", query.CheckInt64IsPositive)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("ArticleController.parseDeleteActionParams"), []byte(err.Error()))
		return -1, err
	}

	return id, nil
}
