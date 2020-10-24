package config

import (
	"Go-Banana/datastruct"
	"Go-Banana/utils"
	"fmt"
)

var version = "0.0.1"

var commConfFile = "config/common_svr.json"

func GetConfVersion() string {
	return version
}

func InitConfig() int {
	svrConfig := datastruct.StructSvrConf{}
	utils.JsonDecode(commConfFile, &svrConfig)
	fmt.Println("Init Config ResVersion:", svrConfig.ResVersion, ", CodeVersion:", svrConfig.CodeVersion)
	return 0
}
