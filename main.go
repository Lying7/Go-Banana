package main

import (
	"Go-Banana/config"
	"Go-Banana/glog"
	"Go-Banana/server"
	"flag"

	"github.com/gin-gonic/gin"
)

func main() {
	isInit := flag.Bool("init", false, "initialize all")
	flag.Parse()
	if *isInit {
		server.InitAll()
	}
	glog.Debug("Go Banana!")

	ptrGinEngine := gin.Default()
	if ptrGinEngine == nil {
		glog.Error("Gin Framework Init Fail!")
		return
	}

	ptrGinEngine.Any("/", DealMsg)
	ptrGinEngine.Run(config.Config().Address)
	return
}
