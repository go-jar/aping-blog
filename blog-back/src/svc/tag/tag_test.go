package tag

import (
	"blog/conf"
	"blog/entity"
	"blog/errno"
	"blog/resource"
	"blog/resource/log"
	"fmt"
	"github.com/go-jar/mysql"
	"os"
	"testing"
	"time"
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

func TestTagSvc(t *testing.T) {
	prepare()

	tagSvc := NewSvc([]byte("traceCategorySvc"))

	ids, err := tagSvc.Insert(
		&entity.TagEntity{TagName: "a1", CreatedTime: time.Now(), UpdatedTime: time.Now()},
		&entity.TagEntity{TagName: "a2", CreatedTime: time.Now(), UpdatedTime: time.Now()},
	)
	t.Log(ids, err)

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

func TestSvc_GetByIds(t *testing.T) {
	prepare()

	tagSvc := NewSvc([]byte("traceCategorySvc"))
	var tagIds []int64
	tagIds = append(tagIds, 12)
	tagIds = append(tagIds, 14)
	tagEntities, err := tagSvc.GetByIds(tagIds)
	if err != nil {
		t.Error(err)
	}

	for _, tag := range tagEntities {
		t.Log(tag.TagName)
	}
}