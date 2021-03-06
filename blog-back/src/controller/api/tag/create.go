package tag

import (
	"time"

	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/entity"
	"blog/errno"
)

func (tc *TagController) CreateAction(context *TagContext) {
	if err := tc.VerifyToken(context.ApiContext); err != nil {
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return
	}

	tagEntity, e := tc.parseCreateActionParams(context)
	if e != nil {
		context.ApiData.Err = e
		return
	}

	ids, err := context.tagSvc.Insert(tagEntity)
	if err != nil {
		context.ApiData.Err = goerror.New(errno.ESysMysqlError, err.Error())
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"TagId":        ids[0],
		"RequestId": context.TraceId,
	}
}

func (tc *TagController) parseCreateActionParams(context *TagContext) (*entity.TagEntity, *goerror.Error) {
	tagEntity := &entity.TagEntity{
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	qs := query.NewQuerySet()
	qs.StringVar(&tagEntity.TagName, "TagName", true, errno.ECommonInvalidArg, "invalid TagName", query.CheckStringNotEmpty)
	qs.Int64Var(&tagEntity.TagIndex, "TagIndex", false, errno.ECommonInvalidArg, "invalid TagIndex", query.CheckInt64GreaterEqual0)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("TagController.parseCreateActionParams"), []byte(err.Error()))
		return nil, err
	}

	return tagEntity, nil
}
