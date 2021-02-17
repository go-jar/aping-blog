package category

import (
	"net/http"

	ctl "github.com/go-jar/gohttp/controller"

	"blog/controller/api"
	"blog/svc/category"
)

type CategoryContext struct {
	*api.ApiContext

	categorySvc *category.Svc
}

func (c *CategoryContext) BeforeAction() {
	c.ApiContext.BeforeAction()

	c.categorySvc = category.NewSvc(c.TraceId)
}

type CategoryController struct {
	api.BaseController
}

func (cc *CategoryController) NewActionContext(req *http.Request, respWriter http.ResponseWriter) ctl.ActionContext {
	ctx := new(CategoryContext)
	ctx.ApiContext = cc.BaseController.NewActionContext(req, respWriter).(*api.ApiContext)

	return ctx
}
