// Package app provide the main application package
package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getIndex(ctx *gin.Context) {
	renderPage("index", ctx)
}

func renderPage(page string, ctx *gin.Context) {
	index, _ := dist.ReadFile(fmt.Sprintf("dist/%s.html", page))
	ctx.Writer.Header().Add("Content-Type", "text/html")
	ctx.Writer.Write(index)
}
