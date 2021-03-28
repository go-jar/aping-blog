package article

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-jar/goerror"
	ctl "github.com/go-jar/gohttp/controller"
	"github.com/go-jar/gohttp/query"

	"blog/controller/api"
	"blog/entity"
	"blog/errno"
	"blog/svc/article"
	"blog/svc/article_tag"
)

type ArticleContext struct {
	*api.ApiContext

	articleSvc    *article.Svc
	articleTagSvc *article_tag.Svc
}

func (ac *ArticleContext) BeforeAction() {
	ac.ApiContext.BeforeAction()

	ac.articleSvc = article.NewSvc(ac.TraceId)
	ac.articleTagSvc = article_tag.NewSvc(ac.TraceId)
}

type ArticleController struct {
	api.BaseController
}

func (ac *ArticleController) NewActionContext(req *http.Request, respWriter http.ResponseWriter) ctl.ActionContext {
	ctx := new(ArticleContext)
	ctx.ApiContext = ac.BaseController.NewActionContext(req, respWriter).(*api.ApiContext)
	ctx.ApiData.Data = map[string]interface{}{
		"RequestId": ctx.TraceId,
	}
	return ctx
}

func (ac *ArticleController) parseCreateOrModifyActionParams(context *ArticleContext, isUpdate bool) (*entity.ArticleEntity, []int64, *goerror.Error) {
	articleEntity := &entity.ArticleEntity{
		ReadCount:   0,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	qs := query.NewQuerySet()
	tagIds := ""
	if isUpdate {
		qs.Int64Var(&articleEntity.Id, "ArticleId", true, errno.ECommonInvalidArg, "invalid Id", query.CheckInt64GreaterEqual0)
	}
	qs.StringVar(&articleEntity.Title, "Title", true, errno.ECommonInvalidArg, "invalid Title", query.CheckStringNotEmpty)
	qs.Int64Var(&articleEntity.CategoryId, "CategoryId", true, errno.ECommonInvalidArg, "invalid CategoryId", query.CheckInt64GreaterEqual0)
	qs.StringVar(&articleEntity.Content, "Content", true, errno.ECommonInvalidArg, "invalid Content", query.CheckStringNotEmpty)
	qs.StringVar(&tagIds, "TagIds", false, errno.ECommonInvalidArg, "invalid TagIds", query.CheckStringNotEmpty)

	if err := qs.Parse(context.QueryValues); err != nil {
		context.ErrorLog([]byte("ArticleController.parseCreateOrModifyActionParams"), []byte(err.Error()))
		return nil, nil, err
	}

	var tagIdsInt64 []int64
	var err *goerror.Error
	if tagIds != "" {
		tagIdsInt64, err = ac.stringToInt64Array(tagIds)
		if err != nil {
			return nil, nil, goerror.New(errno.EInternalError, "TagId: convert string to int64 error")
		}
	}

	return articleEntity, tagIdsInt64, nil
}

func (ac *ArticleController) stringToInt64Array(tagIds string) ([]int64, *goerror.Error) {
	var tagIdsInt64 []int64
	tagSet := make(map[string]bool)
	tagIdList := strings.Split(tagIds, ",")

	for _, tagId := range tagIdList {
		tid, e := strconv.ParseInt(tagId, 10, 64)
		if e != nil {
			return nil, goerror.New(errno.ECommonInvalidArg, "TagId convert error")
		}

		if tid < 0 {
			return nil, goerror.New(errno.ECommonInvalidArg, "TagId negative")
		}

		if _, ok := tagSet[tagId]; ok {
			return nil, goerror.New(errno.ECommonInvalidArg, "TagIds duplicated")
		}

		tagSet[tagId] = true
		tagIdsInt64 = append(tagIdsInt64, tid)
	}

	return tagIdsInt64, nil
}
