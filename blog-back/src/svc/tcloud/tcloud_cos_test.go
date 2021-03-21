package tcloud

import (
	"fmt"
	"os"
	"testing"

	"blog/conf"
	"blog/errno"
	"blog/resource"
	"blog/resource/log"
)

func prepare() {
	if err := conf.Init("/usr/local/lighthouse/softwares/aping-blog/blog-back/"); err != nil {
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
}

func TestTCloudRequest(t *testing.T) {
	prepare()
	ts := NewSvc([]byte("TCloudSvcTest"))

	f, err := os.Open("/data/1.png")
	if err != nil {
		t.Error(err)
	}

	imgUrl, err := ts.PutImg(f)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(imgUrl)
	}
}
