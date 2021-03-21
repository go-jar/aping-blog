package category

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

	categorySvc := NewSvc([]byte("traceCategorySvc"))

	ids, err := categorySvc.Insert(
		&entity.CategoryEntity{CategoryName: "a1", CreatedTime: time.Now(), UpdatedTime: time.Now()},
		&entity.CategoryEntity{CategoryName: "a2", CreatedTime: time.Now(), UpdatedTime: time.Now()},
	)
	t.Log(ids, err)

	item, err := categorySvc.GetById(ids[0])
	t.Log(item, err)

	categoryEntity := entity.CategoryEntity{CategoryName: "b1"}
	if _, err := categorySvc.UpdateById(ids[0], &categoryEntity, map[string]bool{"category_name": true}); err != nil {
		t.Log(err)
	}

	qp := &mysql.QueryParams{
		OrderBy: "id desc",
		Offset:  0,
		Cnt:     10,
	}
	entities, err := categorySvc.SimpleQueryAnd(qp)
	for _, entity := range entities {
		t.Log("SimpleQueryAnd", entity, err)
	}

	total, err := categorySvc.TotalRows(categorySvc.EntityName, categorySvc.RedisKeyPrefix, TotalRowsCacheExpireSeconds)
	t.Log("total:", total, err)

	for _, id := range ids {
		deleted, err := categorySvc.DeleteById(id)
		t.Log(deleted, err)
	}
}
