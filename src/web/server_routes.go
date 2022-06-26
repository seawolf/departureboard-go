package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes(engine *gin.Engine) {
	engine.GET("/ping", handlePing)

	engine.GET("/departure_board/data", handleDepartureBoardData)

	engine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/departure_board/data")
	})
}
