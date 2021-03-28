package entity

import (
	"reflect"
	"time"
)

type CategoryEntity struct {
	Id            int64     `mysql:"id" redis:"id" json:"CategoryId"`
	CategoryName  string    `mysql:"category_name" redis:"category_name" json:"CategoryName"`
	CategoryIndex int64     `mysql:"category_index" redis:"category_index" json:"CategoryIndex"`
	CreatedTime   time.Time `mysql:"created_time" redis:"created_time" json:"CreatedTime"`
	UpdatedTime   time.Time `mysql:"updated_time" redis:"updated_time" json:"UpdatedTime"`
}

var CategoryEntityType = reflect.TypeOf(CategoryEntity{})

type CategoryResultEntity struct {
	Category     *CategoryEntity `json:"Category"`
	ArticleCount int64 	         `json:"ArticleCount"`
}
