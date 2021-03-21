package entity

import (
	"reflect"
	"time"
)

type ArticleTagEntity struct {
	Id          int64     `mysql:"id" redis:"id" json:"Id"`
	ArticleId   int64     `mysql:"article_id" redis:"article_id" json:"ArticleId"`
	TagId       int64     `mysql:"tag_id" redis:"tag_id" json:"TagId"`
	CreatedTime time.Time `mysql:"created_time" redis:"created_time" json:"CreatedTime"`
	UpdatedTime time.Time `mysql:"updated_time" redis:"updated_time" json:"UpdatedTime"`
}

var ArticleTagEntityType = reflect.TypeOf(ArticleTagEntity{})
