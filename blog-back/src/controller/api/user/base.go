package user

import (
	"net/http"

	ctl "github.com/go-jar/gohttp/controller"

	"blog/controller/api"
	"blog/svc/user"
)

type UserContext struct {
	*api.ApiContext

	userSvc *user.Svc
}

func (c *UserContext) BeforeAction() {
	c.ApiContext.BeforeAction()

	c.userSvc = user.NewSvc(c.TraceId)
}

type UserController struct {
	api.BaseController
}

func (uc *UserController) NewActionContext(req *http.Request, respWriter http.ResponseWriter) ctl.ActionContext {
	ctx := new(UserContext)
	ctx.ApiContext = uc.BaseController.NewActionContext(req, respWriter).(*api.ApiContext)

	return ctx
}
