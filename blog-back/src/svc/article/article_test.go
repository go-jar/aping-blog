package article

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-jar/mysql"

	"blog/conf"
	"blog/entity"
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

func TestArticleSvc(t *testing.T) {
	prepare()

	articleSvc := NewSvc([]byte("traceArticleSvc"))

	ids, err := articleSvc.Insert(
		&entity.ArticleEntity{Title: "a1", CategoryId: 1, Content: "sss", ReadCount: 0, CreatedTime: time.Now(), UpdatedTime: time.Now()},
		&entity.ArticleEntity{Title: "a2", CategoryId: 2, Content: "aaa", ReadCount: 0, CreatedTime: time.Now(), UpdatedTime: time.Now()},
	)
	t.Log(ids, err)

	articleEntity := entity.ArticleEntity{Title: "b1"}
	if _, err := articleSvc.UpdateById(ids[0], &articleEntity, map[string]bool{"title": true}); err != nil {
		t.Log(err)
	}

	qp := &mysql.QueryParams{
		OrderBy: "id desc",
		Offset:  0,
		Cnt:     10,
	}
	entities, total, err := articleSvc.GetArticles(qp)
	t.Log("total: ", total)
	for _, item := range entities {
		t.Log("GetArticles", item, err)
	}

	for _, id := range ids {
		deleted, err := articleSvc.DeleteById(id)
		t.Log(deleted, err)
	}
}
