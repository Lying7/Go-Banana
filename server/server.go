package server

import (
	"Go-Banana/config"
	"Go-Banana/glog"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ptrGinEngine *(gin.Engine)

/*Start ...
*
 */
func Start() {
	ptrGinEngine = gin.Default()
	if ptrGinEngine == nil {
		glog.Error("Gin Framework Init Fail!")
		return
	}

	ptrGinEngine.Any("/", DealMsg)
	var listenPort string = fmt.Sprintf(":%d", config.ServerCfg.Port)
	glog.Debug(listenPort)
	ptrGinEngine.Run(listenPort)
	return
}

/*Stop ...
*
 */
func Stop() {
	return
}

/*Tick ...
*
 */
func Tick() {
	for {
		break
	}
	return
}

/*DealMsg ...
*
 */
func DealMsg(context *gin.Context) {
	context.String(http.StatusOK, "Go Banana!")
}
