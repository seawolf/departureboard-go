package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/station"
)

func Prepare(ctx context.Context, client *entities.Client) {
	count, _ := countTOCs(ctx, client)
	if count == 0 {
		createTOCs(ctx, client)
	}

	count, _ = countStations(ctx, client)
	if count == 0 {
		createStations(ctx, client)
	}

	count, _ = countPlatforms(ctx, client)
	if count == 0 {
		createPlatforms(ctx, client)
	}

	count, _ = countDays(ctx, client)
	if count == 0 {
		createDays(ctx, client)
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

func countStations(ctx context.Context, client *entities.Client) (stationCount int, err error) {
	stationCount, err = client.Station.
		Query().
		Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("failed counting Stations: %w", err)
	}

	log.Println("Stations found:", stationCount)
	return
}

func createStations(ctx context.Context, client *entities.Client) {
	count := 0

	client.Station.
		Create().
		SetName("Portsmouth Harbour").
		SetCrs("PMH").
		Save(ctx)
	count++

	client.Station.
		Create().
		SetName("Portsmouth & Southsea").
		SetCrs("PMS").
		Save(ctx)
	count++

	client.Station.
		Create().
		SetName("Fratton").
		SetCrs("FTN").
		Save(ctx)
	count++

	client.Station.
		Create().
		SetName("Hilsea").
		SetCrs("HLS").
		Save(ctx)
	count++

	client.Station.
		Create().
		SetName("Bedhampton").
		SetCrs("BDH").
		Save(ctx)
	count++

	client.Station.
		Create().
		SetName("Havant").
		SetCrs("HAV").
		Save(ctx)
	count++

	log.Println("Stations created:", count)
}

func countPlatforms(ctx context.Context, client *entities.Client) (platformsCount int, err error) {
	platformsCount, err = client.Platform.
		Query().
		Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("failed counting Platforms: %w", err)
	}

	log.Println("Platforms found:", platformsCount)
	return
}

func createPlatforms(ctx context.Context, client *entities.Client) {
	count := 0

	stationPMH, _ := client.Station.
		Query().
		Where(station.Crs("PMH")).
		First(ctx)
	for _, platform := range []string{"1", "2", "3", "4", "5"} {
		client.Platform.
			Create().
			SetStation(stationPMH).
			SetName(platform).
			Save(ctx)
		count++
	}

	stationPMS, _ := client.Station.
		Query().
		Where(station.Crs("PMS")).
		First(ctx)
	for _, platform := range []string{"1", "2", "3", "4"} {
		client.Platform.
			Create().
			SetStation(stationPMS).
			SetName(platform).
			Save(ctx)
		count++
	}

	stationFTN, _ := client.Station.
		Query().
		Where(station.Crs("FTN")).
		First(ctx)
	for _, platform := range []string{"1", "2", "3"} {
		client.Platform.
			Create().
			SetStation(stationFTN).
			SetName(platform).
			Save(ctx)
		count++
	}

	stationHLS, _ := client.Station.
		Query().
		Where(station.Crs("HLS")).
		First(ctx)
	for _, platform := range []string{"1", "2"} {
		client.Platform.
			Create().
			SetStation(stationHLS).
			SetName(platform).
			Save(ctx)
		count++
	}

	stationBDH, _ := client.Station.
		Query().
		Where(station.Crs("BDH")).
		First(ctx)
	for _, platform := range []string{"1", "2"} {
		client.Platform.
			Create().
			SetStation(stationBDH).
			SetName(platform).
			Save(ctx)
		count++
	}

	stationHAV, _ := client.Station.
		Query().
		Where(station.Crs("HAV")).
		First(ctx)
	for _, platform := range []string{"1", "2"} {
		client.Platform.
			Create().
			SetStation(stationHAV).
			SetName(platform).
			Save(ctx)
		count++
	}

	log.Println("Platforms created:", count)
}

func countDays(ctx context.Context, client *entities.Client) (daysCount int, err error) {
	daysCount, err = client.Day.
		Query().
		Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("failed counting Days: %w", err)
	}

	log.Println("Days found:", daysCount)
	return
}

func createDays(ctx context.Context, client *entities.Client) {
	count := 0

	for _, numberOfDays := range []int{6, 5, 4, 3, 2, 1, 0} {

		now := time.Now()
		startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

		newDay := startOfToday.AddDate(0, 0, -numberOfDays)

		client.Day.
			Create().
			SetDate(newDay).
			Save(ctx)
		count++
	}

	log.Println("Days created:", count)
}
