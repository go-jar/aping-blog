package entity

type UserEntity struct {
	Id          int64  `mysql:"id" redis:"id" json:"Id"`
	Username    string `mysql:"username" redis:"username" json:"Username"`
	Password    string `mysql:"password" redis:"password" json:"Password"`
	Role        string `mysql:"role" redis:"role" json:"Role"`
	CreatedTime string `mysql:"created_time" redis:"created_time" json:"CreatedTime"`
	UpdatedTime string `mysql:"updated_time" redis:"updated_time" json:"UpdatedTime" `
}
