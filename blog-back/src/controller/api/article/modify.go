package article

import (
	"github.com/go-jar/goerror"

	"blog/errno"
)

func (ac *ArticleController) ModifyAction(context *ArticleContext) {
	if err := ac.VerifyToken(context.ApiContext); err != nil {
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return
	}

	articleEntity, newTagIds, err := ac.parseCreateOrModifyActionParams(context, true)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	required := map[string]bool{
		"title":       true,
		"category_id": true,
		"content":     true,
	}
	if _, err := context.articleSvc.UpdateById(articleEntity.Id, articleEntity, required); err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	if err := context.articleTagSvc.UpdateArticleTags(articleEntity.Id, newTagIds); err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"RequestId": context.TraceId,
	}
}
