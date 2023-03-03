package util

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"facette.io/natsort"
	"github.com/gin-gonic/gin"
)

func ListFile(root string) ([]string, error) {
	pattern := root + "/*"
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	for i := range matches {
		news := strings.Replace(
			filepath.Base(matches[i]),
			filepath.Base(root), "", 1)
		if info, err := os.Stat(matches[i]); err == nil && info.IsDir() {
			news = news + "/"
		}
		matches[i] = news
	}
	natsort.Sort(matches)
	return matches, nil
}

func GinBasic(router *gin.Engine, root string) {
	router.LoadHTMLGlob("src/html/**/*")
	router.GET("/file/*path", func(ctx *gin.Context) {
		path := ctx.Param("path")
		fullpath := root + path
		println(fullpath)
		info, err := os.Stat(fullpath)
		if err != nil {
			println(fullpath)
			println(err)
			ctx.String(404, "not found")
			return
		}
		if info.IsDir() {
			list, _ := ListFile(fullpath)
			ctx.HTML(200, "index.html", gin.H{
				"path":    path,
				"list":    list,
				"notRoot": path != "/",
			})
		} else {
			file, err := os.Open(fullpath)
			if err != nil {
				ctx.String(500, "Some error happen.")
			}
			http.ServeContent(ctx.Writer, ctx.Request, info.Name(), info.ModTime(), file)
		}
		// ctx.JSON(200, gin.H{
		// 	"hello":    "gin",
		// 	"root":     root,
		// 	"path":     path,
		// 	"fullpath": fullpath,
		// 	"info": gin.H{
		// 		"name":  info.Name(),
		// 		"IsDir": info.IsDir(),
		// 	},
		// })

	})
}
