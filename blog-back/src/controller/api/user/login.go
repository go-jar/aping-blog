package user

import (
	"github.com/go-jar/goerror"
	"github.com/go-jar/gohttp/query"

	"blog/conf"
	"blog/entity"
	"blog/errno"
	"blog/utils"
)

func (uc *UserController) LoginAction(context *UserContext) {
	userFromRequest, e := uc.parseLoginActionParams(context)
	if e != nil {
		context.ApiData.Err = e
		return
	}

	userFromDb, err := context.userSvc.GetByUsername(userFromRequest.Username)
	if err != nil {
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, err.Error())
		return
	}

	if !context.userSvc.VerifyPassword(userFromRequest.Password, userFromDb.Password) {
		context.ErrorLog([]byte("UserController.LoginAction"), []byte("invalid password"))
		context.ApiData.Err = goerror.New(errno.EUserUnauthorized, "invalid password")
		return
	}

	token, err := utils.GenerateToken(userFromDb.Id, userFromDb.Username, conf.LoginSecretKey, conf.LoginExpireSeconds)
	if err != nil {
		context.ApiData.Err = e
		return
	}

	context.ApiData.Data = map[string]interface{}{
		"User":  userFromDb,
		"Token": token,
		"RequestId": context.TraceId,
	}
}

func (uc *UserController) parseLoginActionParams(context *UserContext) (*entity.UserEntity, *goerror.Error) {
	userEntity := new(entity.UserEntity)

	qs := query.NewQuerySet()
	qs.StringVar(&userEntity.Username, "Username", true, errno.ECommonInvalidArg, "invalid Username", query.CheckStringNotEmpty)
	qs.StringVar(&userEntity.Password, "Password", true, errno.ECommonInvalidArg, "invalid Password", query.CheckStringNotEmpty)

	if e := qs.Parse(context.QueryValues); e != nil {
		context.ErrorLog([]byte("UserController.parseLoginActionParams"), []byte(e.Error()))
		return nil, e
	}

	return userEntity, nil
}
