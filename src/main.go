package main

import (
	"context"

	"bitbucket.org/sea_wolf/departure_board-go/v2/db"
	"bitbucket.org/sea_wolf/departure_board-go/v2/web"
)

func main() {
	ctx := context.Background()

	client := db.Connect(ctx)
	defer client.Close()

	db.Prepare(ctx, client)

	web.StartServer(ctx, client)
}
