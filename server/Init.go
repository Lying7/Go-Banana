package server

import (
	"Go-Banana/config"
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
)

func InitAll() {
	InitDB()
	InitGin()
}

func InitDB() {
	return
}

func InitGin() {
	router = gin.Default()
	logFilePath := config.Config().LogPath
	logFileName := config.Config().LogName

	fileName := path.Join(logFilePath, logFileName)
	fmt.Println("path", fileName)
	//router.Use(middleware.LoggerToFile())
}
