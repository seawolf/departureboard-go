package web

import "github.com/gin-gonic/gin"

func views(engine *gin.Engine) {
	engine.LoadHTMLGlob("web/pages/**/*.tmpl")
}
