package api

import (
	"html"
	"net/http"

	"github.com/go-jar/goerror"
	ctl "github.com/go-jar/gohttp/controller"

	"blog/controller"
	"blog/utils"
)

type IApiDataContext interface {
	ctl.ActionContext

	Version() string
	Data() interface{}
	Err() *goerror.Error
}

type ApiContext struct {
	*controller.BaseContext

	ApiData struct {
		Version string
		Data interface{}
		Err  *goerror.Error
	}
}

func (ac *ApiContext) Version() string {
	return ac.ApiData.Version
}

func (ac *ApiContext) Data() interface{} {
	return ac.ApiData.Data
}

func (ac *ApiContext) Err() *goerror.Error {
	return ac.ApiData.Err
}

func (ac *ApiContext) AfterAction() {
	f := ac.QueryValues.Get("fmt")
	if f == "jsonp" {
		callback := ac.QueryValues.Get("_callback")
		if callback != "" {
			ac.SetResponseBody(utils.ApiJsonp(ac.ApiData.Version, ac.ApiData.Data, ac.ApiData.Err, html.EscapeString(callback)))
			return
		}
	}

	ac.SetResponseBody(utils.ApiJson(ac.ApiData.Version, ac.ApiData.Data, ac.ApiData.Err))
	ac.BaseContext.AfterAction()
}

type BaseController struct {
	controller.BaseController
}

func (bc *BaseController) NewActionContext(req *http.Request, respWriter http.ResponseWriter) ctl.ActionContext {
	context := new(ApiContext)
	context.BaseContext = bc.BaseController.NewActionContext(req, respWriter).(*controller.BaseContext)

	return context
}

func JumpToApiError(context ctl.ActionContext, args ...interface{}) {
	acontext := context.(IApiDataContext)

	acontext.SetResponseBody(utils.ApiJson(acontext.Version(), acontext.Data(), acontext.Err()))
}