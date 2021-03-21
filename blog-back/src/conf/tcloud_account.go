package conf

import (
	"blog/utils"
)

type tcloudAccountJson struct {
	AppId         string `json:"app_id"`
	Uin           string `json:"uin"`
	SubAccountUin string `json:"sub_account_uin"`
	SecretId      string `json:"secret_id"`
	SecretKey     string `json:"secret_key"`
	CosHost		  string `json:"cos_host"`
}

var TcloudAccount struct {
	AppId         string
	Uin           string
	SubAccountUin string
	SecretId      string
	SecretKey     string
	CosHost 	  string
}

func initTcloudAccountConf() {
	TcloudAccount.AppId = serverConf.TcloudAccount.AppId
	TcloudAccount.Uin = serverConf.TcloudAccount.Uin
	TcloudAccount.SubAccountUin = serverConf.TcloudAccount.SubAccountUin
	TcloudAccount.SecretId = utils.DecryptString(serverConf.TcloudAccount.SecretId)
	TcloudAccount.SecretKey = utils.DecryptString(serverConf.TcloudAccount.SecretKey)
	TcloudAccount.CosHost = serverConf.TcloudAccount.CosHost
}
