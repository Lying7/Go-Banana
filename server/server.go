package server

import (
	"Go-Banana/config"
	"Go-Banana/glog"
	"flag"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	isInit := flag.Bool("init", false, "initialize all")
	flag.Parse()
	if *isInit {
		InitAll()
	}
	glog.Debug("Go Banana!")

	ptrGinEngine := gin.Default()
	if ptrGinEngine == nil {
		glog.Error("Gin Framework Init Fail!")
		return
	}

	if err := ptrGinEngine.Run(config.Config().Address); err != nil {
		os.Exit(1)
	}
	return
}
