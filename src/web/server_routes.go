package web

import (
	"github.com/gin-gonic/gin"
)

func routes(engine *gin.Engine) {
	engine.GET("/ping", handlePing)
}
