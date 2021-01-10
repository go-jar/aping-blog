package entity

type UserEntity struct {
	Id          int64    `mysql:"id" json:"id" redis:"id"`
	Username    string `mysql:"username" json:"username" redis:"username"`
	Password    string `mysql:"password" json:"password" redis:"password"`
	Role        string `mysql:"role" json:"role" redis:"role"`
	CreatedTime string `mysql:"created_time" json:"created_time" redis:"created_time"`
	UpdatedTime string `mysql:"updated_time" json:"updated_time" redis:"updated_time"`
}
