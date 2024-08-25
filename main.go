package main

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin"
	"github.com/loebfly/ezgin/app"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/ezlogs"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"net/http"
	"soraka/define"
	"soraka/schedule"
)

// @title soraka API
// @description 对MAC系统下的文件简单操作的API
// @version 1.0.0
func main() {
	ezlogs.CDebug("soraka API", "对MAC系统下的文件简单操作的API")

	ezgin.Start(app.Start{
		GinCfg: app.GinCfg{
			RecoveryHandler: func(c *gin.Context, err any) {
				c.JSON(http.StatusOK, engine.ErrorRes(-1, "系统异常"))
			},
			NoRouteHandler: func(c *gin.Context) {
				c.JSON(http.StatusOK, engine.ErrorRes(-1, "接口不存在"))
			},
		},
	})

	define.Setup()

	schedule.Start()

	ezgin.ShutdownWhenExitSignal()
}
