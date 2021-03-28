package entity

import (
	"reflect"
	"time"
)

type RemarkEntity struct {
	Id              int64     `mysql:"id" redis:"id" json:"RemarkId"`
	ArticleId  	    int64     `mysql:"article_id" redis:"article_id" json:"ArticleId"`
	Nickname 	    string    `mysql:"nickname" redis:"nickname" json:"Nickname"`
	Content 	    string    `mysql:"content" redis:"content" json:"Content"`
	InitRemarkId    int64     `mysql:"init_remark_id" redis:"init_remark_id" json:"InitRemarkId"`
	NicknameReplied string    `mysql:"nickname_replied" redis:"nickname_replied" json:"NicknameReplied"`
	CreatedTime     time.Time `mysql:"created_time" redis:"created_time" json:"CreatedTime"`
	UpdatedTime     time.Time `mysql:"updated_time" redis:"updated_time" json:"UpdatedTime"`
}

var RemarkEntityType = reflect.TypeOf(RemarkEntity{})

type RemarkResultEntity struct {
	Remark   *RemarkEntity   `json:"Remark"`
	ReplySet []*RemarkEntity `json:"ReplySet"`
}
