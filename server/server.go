package server

import (
	"Go-Banana/config"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Start() {
	// isInit := flag.Bool("init", false, "initialize all")
	// flag.Parse()
	// if *isInit {
	// 	InitAll()
	// }
	InitAll()

	if err := router.Run(config.Config().Address); err != nil {
		os.Exit(1)
	}
}
