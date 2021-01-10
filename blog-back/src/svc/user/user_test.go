package user

import (
	"blog/conf"
	"testing"

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
	_ = conf.Init("/data/blog-back/")
	resource.InitMysql()
	resource.InitRedis()
}