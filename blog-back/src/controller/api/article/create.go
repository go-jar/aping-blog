package article

import (
	"github.com/go-jar/goerror"

	"blog/errno"
)

func (ac *ArticleController) CreateAction(context *ArticleContext) {
	articleEntity, tagIds, e := ac.parseCreateOrUpdateActionParams(context, false)
	if e != nil {
		context.ApiData.Err = e
		return
	}

	ids, err := context.articleSvc.Insert(articleEntity)
	if err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	if len(tagIds) > 0 {
		if _, err := context.articleTagSvc.Insert(ids[0], tagIds...); err != nil {
			context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
			return
		}
	}

	context.ApiData.Data = map[string]interface{}{
		"ArticleId":        ids[0],
		"RequestId": context.TraceId,
	}
}
