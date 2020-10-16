package main

import (
	"Go-Banana/conf"    //config
	"Go-Banana/dbProxy" //db proxy
	"Go-Banana/utils"   //utils

	//utils and ti
	"github.com/gin-gonic/gin" //gin framework
)

func main() {
	ginVer := gin.Version
	dbProxyVer := dbProxy.GetDBProxyVersion()
	confVer := conf.GetConfVersion()
	utilsVer := utils.GetUtilsVersion()
	println("Simple demo with gin[", ginVer, "], dbProxy[", dbProxyVer, "], conf[", confVer, "], utils[", utilsVer, "]")
}
