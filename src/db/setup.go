package db

import (
	"context"
	"log"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(ctx context.Context) (client *entities.Client) {
	client, err := entities.Open("sqlite3", "file:../db/dev.sqlite3?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return
}
