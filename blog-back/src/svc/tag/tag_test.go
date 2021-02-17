package tag

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

func TestCategorySvc(t *testing.T) {
	prepare()

	tagSvc := NewSvc([]byte("traceCategorySvc"))

	ids, err := tagSvc.Insert(
		&entity.TagEntity{TagName: "a1", CreatedTime: time.Now(), UpdatedTime: time.Now()},
		&entity.TagEntity{TagName: "a2", CreatedTime: time.Now(), UpdatedTime: time.Now()},
	)
	t.Log(ids, err)

	item, err := tagSvc.GetById(ids[0])
	t.Log(item, err)

	tagEntity := entity.TagEntity{TagName: "b1"}
	if _, err := tagSvc.UpdateById(ids[0], &tagEntity, map[string]bool{"tag_name": true}); err != nil {
		t.Log(err)
	}

	qp := &mysql.QueryParams{
		OrderBy: "id desc",
		Offset:  0,
		Cnt:     10,
	}
	entities, err := tagSvc.SimpleQueryAnd(qp)
	for _, entity := range entities {
		t.Log("SimpleQueryAnd", entity, err)
	}

	total, err := tagSvc.SimpleTotalAnd(qp)
	t.Log("total:", total, err)

	for _, id := range ids {
		deleted, err := tagSvc.DeleteById(id)
		t.Log(deleted, err)
	}
}
