package main

import (
	"context"
	"fmt"
	"log"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/toc"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()

	client := setupDatabase(ctx)
	defer client.Close()

	count, _ := countTOCs(ctx, client)
	if count == 0 {
		createTOCs(ctx, client)
	}

	getTOCs(ctx, client)
}

func setupDatabase(ctx context.Context) (client *entities.Client) {
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

func countTOCs(ctx context.Context, client *entities.Client) (tocCount int, err error) {
	tocCount, err = query(client).
		Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("failed counting TOCs: %w", err)
	}

	log.Println("TOCs found:", tocCount)
	return
}

func createTOCs(ctx context.Context, client *entities.Client) {
	count := 0

	client.TOC.
		Create().
		SetName("South Western Railway").
		Save(ctx)
	count++

	client.TOC.
		Create().
		SetName("Southern").
		Save(ctx)
	count++

	log.Println("TOCs created:", count)
}

func getTOCs(ctx context.Context, client *entities.Client) (tocs []*entities.TOC, err error) {
	tocs, err = query(client).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying TOCs: %w", err)
	}

	log.Println("TOCs returned:", tocs)
	return
}

func query(client *entities.Client) *entities.TOCQuery {
	return client.TOC.
		Query().
		Where(toc.NameContains("South"))
}
