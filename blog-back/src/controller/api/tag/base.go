package tag

import (
	"net/http"

	ctl "github.com/go-jar/gohttp/controller"

	"blog/controller/api"
	"blog/svc/tag"
)

type TagContext struct {
	*api.ApiContext

	tagSvc *tag.Svc
}

func (c *TagContext) BeforeAction() {
	c.ApiContext.BeforeAction()

	c.tagSvc = tag.NewSvc(c.TraceId)
}

type TagController struct {
	api.BaseController
}

func (tc *TagController) NewActionContext(req *http.Request, respWriter http.ResponseWriter) ctl.ActionContext {
	ctx := new(TagContext)
	ctx.ApiContext = tc.BaseController.NewActionContext(req, respWriter).(*api.ApiContext)

	return ctx
}
