package main

import (
	"blog/controller/api/file"
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"

	"github.com/go-jar/gohttp/gracehttp"
	"github.com/go-jar/gohttp/router"
	"github.com/go-jar/gohttp/system"
	"github.com/go-jar/pidfile"

	"blog/conf"
	"blog/controller/api/article"
	"blog/controller/api/category"
	"blog/controller/api/tag"
	"blog/controller/api/user"
	"blog/errno"
	"blog/resource"
	"blog/resource/log"
)

func main() {
	var prjHome string

	flag.StringVar(&prjHome, "prj-home", "", "prj-home absolute path")
	flag.Parse()

	prjHome = strings.TrimRight(prjHome, "/")
	if prjHome == "" {
		fmt.Println("missing flag prjHome: ")
		flag.PrintDefaults()
		os.Exit(errno.ESysInvalidPrjHome)
	}

	e := conf.Init(prjHome)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(errno.ESysInitConfFail)
	}

	if conf.PprofConf.Enable {
		go func() {
			_ = http.ListenAndServe("127.0.0.1:"+conf.PprofConf.Port, nil)
		}()
	}

	e = log.InitLog("api")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(errno.ESysInitLogFail)
	}
	defer func() {
		log.FreeLog()
	}()

	resource.InitMysql()
	resource.InitRedis()

	pf, err := pidfile.CreatePidFile(conf.CommonConf.ApiPidFile)
	if err != nil {
		fmt.Printf("create pid file %s failed, error: %s\n", conf.CommonConf.ApiPidFile, err.Error())
		os.Exit(errno.ESysSavePidFileFail)
	}

	r := router.NewSimpleRouter(log.TestLogger)
	r.RegisterRoutes(
		new(user.UserController),
		new(category.CategoryController),
		new(tag.TagController),
		new(article.ArticleController),
		new(file.FileController),
	)

	sys := system.NewSystem(r)

	fmt.Printf("listen at %s:%s\n", conf.ApiHttpConf.GoHttpHost, conf.ApiHttpConf.GoHttpPort)

	err = gracehttp.ListenAndServe(conf.ApiHttpConf.GoHttpHost+":"+conf.ApiHttpConf.GoHttpPort, sys)
	if err != nil {
		fmt.Println("pid:" + strconv.Itoa(os.Getpid()) + ", err:" + err.Error())
	}

	if err := pidfile.ClearPidFile(pf); err != nil {
		fmt.Printf("clear pid file failed, error: %s\n", err.Error())
	}
}
