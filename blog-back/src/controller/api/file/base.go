package file

import (
	"net/http"

	ctl "github.com/go-jar/gohttp/controller"

	"blog/controller/api"
	"blog/svc/tcloud"
)

type FileContext struct {
	*api.ApiContext

	tCloudSvc *tcloud.Svc
}

func (ac *FileContext) BeforeAction() {
	ac.ApiContext.BeforeAction()

	ac.tCloudSvc = tcloud.NewSvc(ac.TraceId)
}

type FileController struct {
	api.BaseController
}

func (fc *FileController) NewActionContext(req *http.Request, respWriter http.ResponseWriter) ctl.ActionContext {
	ctx := new(FileContext)
	ctx.ApiContext = fc.BaseController.NewActionContext(req, respWriter).(*api.ApiContext)
	ctx.ApiData.Data = map[string]interface{}{
		"RequestId": ctx.TraceId,
	}
	return ctx
}
