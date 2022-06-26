package web

import (
	"github.com/gin-gonic/gin"
)

func routes(engine *gin.Engine) {
	engine.GET("/ping", handlePing)

	engine.GET("/departure_board/data", handleDepartureBoardData)
}
