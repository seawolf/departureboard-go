package db

import (
	"context"
	"fmt"
	"log"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities"
)

func Prepare(ctx context.Context, client *entities.Client) {
	count, _ := countTOCs(ctx, client)
	if count == 0 {
		createTOCs(ctx, client)
	}
}

func countTOCs(ctx context.Context, client *entities.Client) (tocCount int, err error) {
	tocCount, err = client.TOC.
		Query().
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
