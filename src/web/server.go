package web

import (
	"context"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities"
	"github.com/gin-gonic/gin"
)

var CTX context.Context
var Client *entities.Client

func StartServer(ctx context.Context, client *entities.Client) {
	CTX = ctx
	Client = client

	engine := gin.Default()

	routes(engine)

	engine.Run()
}
