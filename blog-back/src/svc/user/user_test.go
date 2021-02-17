package user

import (
	"blog/errno"
	"blog/resource/log"
	"fmt"
	"os"
	"testing"

	"blog/conf"
	"blog/resource"
)

func TestSvc(t *testing.T) {
	prepare()
	userSvc := NewSvc([]byte("-"))

	userEntity, err := userSvc.GetByUsername("felicity")
	if err != nil {
		t.Error(err)
	}
	t.Log(userEntity)

	if !userSvc.VerifyPassword("123456", userEntity.Password) {
		t.Error("invalid password")
	}
}

func prepare() {
	if err := conf.Init("/data/blog-back/"); err != nil {
		fmt.Println(err)
	}

	err := log.InitLog("api")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(errno.ESysInitLogFail)
	}
	defer func() {
		log.FreeLog()
	}()

	resource.InitMysql()
	resource.InitRedis()
}
