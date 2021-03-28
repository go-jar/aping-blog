package api

import (
	"errors"
	"html"
	"net/http"

	"github.com/go-jar/goerror"
	ctl "github.com/go-jar/gohttp/controller"

	"blog/conf"
	"blog/controller"
	"blog/errno"
	"blog/svc/user"
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
		Data    interface{}
		Err     *goerror.Error
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

func (bc *BaseController) VerifyToken(context *ApiContext) error {
	token :=  context.Request().Header.Get("X-Token")
	userClaim, err := utils.ParseToken(token, conf.LoginSecretKey)
	if err != nil {
		err = errors.New("login expired")
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return err
	}

	userSvc := user.NewSvc(context.TraceId)
	user, err := userSvc.GetByUsername(conf.AdminName)
	if err != nil {
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return err
	}

	if userClaim.Id != user.Id || userClaim.Username != user.Username {
		err = errors.New("invalid user")
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return err
	}

	return nil
}

func JumpToApiError(context ctl.ActionContext, args ...interface{}) {
	acontext := context.(IApiDataContext)

	acontext.SetResponseBody(utils.ApiJson(acontext.Version(), acontext.Data(), acontext.Err()))
}

func CheckInt64GreaterEqual0(v int64) bool {
	if v >= 0 {
		return true
	}
	return false
}
