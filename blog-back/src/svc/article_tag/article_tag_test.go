package article_tag

import (
	"fmt"
	"github.com/go-jar/mysql"
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
	resource.InitRedis()
}

func TestArticleTagSvc(t *testing.T) {
	prepare()
	articleTagSvc := NewSvc([]byte("traceArticleTagSvc"))

	ids, err := articleTagSvc.Insert(0, 1, 2, 3)
	t.Log(ids, err)

	qp := &mysql.QueryParams{
		OrderBy: "id desc",
		Offset:  0,
		Cnt:     10,
	}
	entities, err := articleTagSvc.SimpleQueryAnd(qp)
	for _, item := range entities {
		t.Log("SimpleQueryAnd", item, err)
	}
}

func TestSvc_DeleteByArticleId(t *testing.T) {
	prepare()
	articleTagSvc := NewSvc([]byte("traceArticleTagSvc"))

	if err := articleTagSvc.DeleteByArticleId(23); err != nil {
		t.Error(err)
	} else {
		t.Log("deleted")
	}
}
