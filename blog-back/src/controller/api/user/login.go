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
	userFromRequest, e := uc.parseAddActionParams(context)
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
		"user": userFromDb,
		"token": token,
	}
}

func (uc *UserController) parseAddActionParams(context *UserContext) (*entity.UserEntity, *goerror.Error) {
	userEntity := new(entity.UserEntity)

	qs := query.NewQuerySet()
	qs.StringVar(&userEntity.Username, "username", true, errno.ECommonInvalidArg, "invalid username", query.CheckStringNotEmpty)
	qs.StringVar(&userEntity.Password, "password", true, errno.ECommonInvalidArg, "invalid password", query.CheckStringNotEmpty)

	if e := qs.Parse(context.QueryValues); e != nil {
		context.ErrorLog([]byte("UserController.parseAddActionParams"), []byte(e.Error()))
		return nil, e
	}

	return userEntity, nil
}
