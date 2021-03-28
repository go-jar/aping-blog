package entity

import (
	"reflect"
	"time"
)

type ArticleEntity struct {
	Id          int64     `mysql:"id" redis:"id" json:"ArticleId"`
	Title       string    `mysql:"title" redis:"title" json:"Title"`
	CategoryId  int64     `mysql:"category_id" redis:"category_id" json:"CategoryId"`
	Content     string    `mysql:"content" redis:"content" json:"Content"`
	ReadCount   int64     `mysql:"read_count" redis:"read_count" json:"ReadCount"`
	CreatedTime time.Time `mysql:"created_time" redis:"created_time" json:"CreatedTime"`
	UpdatedTime time.Time `mysql:"updated_time" redis:"updated_time" json:"UpdatedTime"`
}

var ArticleEntityType = reflect.TypeOf(ArticleEntity{})

type ArticleResultEntity struct {
	Article     *ArticleEntity  `json:"Article"`
	Category    *CategoryEntity `json:"Category"`
	TagSet      []*TagEntity    `json:"TagSet"`
	RemarkCount int64           `json:"RemarkCount"`
}
