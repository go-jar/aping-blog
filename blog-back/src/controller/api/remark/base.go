package remark

import (
	"net/http"

	ctl "github.com/go-jar/gohttp/controller"

	"blog/controller/api"
	"blog/svc/remark"
)

type RemarkContext struct {
	*api.ApiContext

	remarkSvc *remark.Svc
}

func (c *RemarkContext) BeforeAction() {
	c.ApiContext.BeforeAction()

	c.remarkSvc = remark.NewSvc(c.TraceId)
}

type RemarkController struct {
	api.BaseController
}

func (rc *RemarkController) NewActionContext(req *http.Request, respWriter http.ResponseWriter) ctl.ActionContext {
	ctx := new(RemarkContext)
	ctx.ApiContext = rc.BaseController.NewActionContext(req, respWriter).(*api.ApiContext)
	ctx.ApiData.Data = map[string]interface{}{
		"RequestId": ctx.TraceId,
	}
	return ctx
}
