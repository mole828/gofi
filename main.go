package main

import (
	"gofi/src/util"

	"github.com/gin-gonic/gin"
)

func main() {
	root := "D:/"
	app := gin.New()
	util.GinBasic(app, root)
	app.Run(":80")
}
