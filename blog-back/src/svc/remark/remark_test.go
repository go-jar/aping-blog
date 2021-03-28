package remark

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

func TestRemarkSvc(t *testing.T) {
	prepare()

	remarkSvc := NewSvc([]byte("traceRemarkSvc"))

	ids, err := remarkSvc.Insert(
		&entity.RemarkEntity{
			ArticleId: 1,
			Nickname: "test",
			Content: "test",
			InitRemarkId: -1,
			NicknameReplied: "",
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		},
		&entity.RemarkEntity{
			ArticleId: 1,
			Nickname: "test1",
			Content: "test1",
			InitRemarkId: 1,
			NicknameReplied: "test",
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		},
	)
	t.Log(ids, err)

	remarkEntity := entity.RemarkEntity{Content: "b1"}
	if _, err := remarkSvc.UpdateById(ids[0], &remarkEntity, map[string]bool{"content": true}); err != nil {
		t.Log(err)
	}

	qp := &mysql.QueryParams{
		OrderBy: "id desc",
		Offset:  0,
		Cnt:     10,
	}
	entities, total, err := remarkSvc.GetRemarks(qp)
	t.Log("total: ", total)
	for _, item := range entities {
		t.Log("GetRemarks", item, err)
	}

	for _, id := range ids {
		deleted, err := remarkSvc.DeleteById(id)
		t.Log(deleted, err)
	}
}
