package tag

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/entity"
	"blog/errno"
)

func (tc *TagController) ModifyAction(context *TagContext) {
	if err := tc.VerifyToken(context.ApiContext); err != nil {
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return
	}

	tagEntity, err := tc.parseModifyActionParams(context)
	if err != nil {
		context.ApiData.Err = err
		return
	}

	required := map[string]bool{
		"tag_name":  true,
		"tag_index": true,
	}
	if _, err := context.tagSvc.UpdateById(tagEntity.Id, tagEntity, required); err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"RequestId": context.TraceId,
	}
}

func (tc *TagController) parseModifyActionParams(context *TagContext) (*entity.TagEntity, *goerror.Error) {
	tagEntity := &entity.TagEntity{}

	qs := query.NewQuerySet()
	qs.Int64Var(&tagEntity.Id, "TagId", true, errno.ECommonInvalidArg, "invalid TagId", query.CheckInt64GreaterEqual0)
	qs.StringVar(&tagEntity.TagName, "TagName", true, errno.ECommonInvalidArg, "invalid TagName", query.CheckStringNotEmpty)
	qs.Int64Var(&tagEntity.TagIndex, "TagIndex", false, errno.ECommonInvalidArg, "invalid TagIndex", query.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("TagController.parseModifyActionParams"), []byte(err.Error()))
		return nil, err
	}

	return tagEntity, nil
}
