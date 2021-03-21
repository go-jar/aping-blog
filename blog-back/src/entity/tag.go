package entity

import (
	"reflect"
	"time"
)

type TagEntity struct {
	Id          int64     `mysql:"id" redis:"id" json:"Id"`
	TagName     string    `mysql:"tag_name" redis:"tag_name" json:"TagName"`
	TagIndex    int64     `mysql:"tag_index" redis:"tag_index" json:"TagIndex"`
	CreatedTime time.Time `mysql:"created_time" redis:"created_time" json:"CreatedTime"`
	UpdatedTime time.Time `mysql:"updated_time" redis:"updated_time" json:"UpdatedTime"`
}

var TagEntityType = reflect.TypeOf(TagEntity{})
