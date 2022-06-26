package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/platform"
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

	count, _ = countCallingPoints(ctx, client)
	if count == 0 {
		createCallingPoints(ctx, client)
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
		SaveX(ctx)
	count++

	client.TOC.
		Create().
		SetName("Southern").
		SaveX(ctx)
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
		SaveX(ctx)
	count++

	client.Station.
		Create().
		SetName("Portsmouth & Southsea").
		SetCrs("PMS").
		SaveX(ctx)
	count++

	client.Station.
		Create().
		SetName("Fratton").
		SetCrs("FTN").
		SaveX(ctx)
	count++

	client.Station.
		Create().
		SetName("Hilsea").
		SetCrs("HLS").
		SaveX(ctx)
	count++

	client.Station.
		Create().
		SetName("Bedhampton").
		SetCrs("BDH").
		SaveX(ctx)
	count++

	client.Station.
		Create().
		SetName("Havant").
		SetCrs("HAV").
		SaveX(ctx)
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
			SaveX(ctx)
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
			SaveX(ctx)
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
			SaveX(ctx)
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
			SaveX(ctx)
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
			SaveX(ctx)
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
			SaveX(ctx)
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
			SaveX(ctx)
		count++
	}

	log.Println("Days created:", count)
}

func countCallingPoints(ctx context.Context, client *entities.Client) (platformsCount int, err error) {
	platformsCount, err = client.CallingPoint.
		Query().
		Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("failed counting CallingPoints: %w", err)
	}

	log.Println("CallingPoints found:", platformsCount)
	return
}

func createCallingPoints(ctx context.Context, client *entities.Client) {
	count := 0

	days, _ := client.Day.
		Query().
		Order(entities.Desc("id")).
		All(ctx)

	PMH_3, _ := client.Platform.
		Query().
		Where(platform.HasStationWith(station.Crs("PMH"))).
		Where(platform.Name("3")).
		First(ctx)

	PMH_4, _ := client.Platform.
		Query().
		Where(platform.HasStationWith(station.Crs("PMH"))).
		Where(platform.Name("4")).
		First(ctx)

	PMS_1, _ := client.Platform.
		Query().
		Where(platform.HasStationWith(station.Crs("PMS"))).
		Where(platform.Name("1")).
		First(ctx)

	FTN_1, _ := client.Platform.
		Query().
		Where(platform.HasStationWith(station.Crs("FTN"))).
		Where(platform.Name("1")).
		First(ctx)

	HLS_1, _ := client.Platform.
		Query().
		Where(platform.HasStationWith(station.Crs("HLS"))).
		Where(platform.Name("1")).
		First(ctx)

	BDH_1, _ := client.Platform.
		Query().
		Where(platform.HasStationWith(station.Crs("BDH"))).
		Where(platform.Name("1")).
		First(ctx)

	HAV_1, _ := client.Platform.
		Query().
		Where(platform.HasStationWith(station.Crs("HAV"))).
		Where(platform.Name("1")).
		First(ctx)

	// 1P24
	client.CallingPoint.
		Create().
		SetPlatform(PMH_4).
		SetArrivalTime(time.Time{}).
		SetDepartureTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 15)).
		SaveX(ctx)
	count++
	client.CallingPoint.
		Create().
		SetPlatform(PMS_1).
		SetArrivalTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 18)).
		SetDepartureTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 20)).
		SaveX(ctx)
	count++
	client.CallingPoint.
		Create().
		SetPlatform(FTN_1).
		SetArrivalTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 23)).
		SetDepartureTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 24)).
		SaveX(ctx)
	count++
	client.CallingPoint.
		Create().
		SetPlatform(HAV_1).
		SetArrivalTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 32)).
		SetDepartureTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 34)).
		SaveX(ctx)
	count++

	// 2P98
	client.CallingPoint.
		Create().
		SetPlatform(PMH_3).
		SetArrivalTime(time.Time{}).
		SetDepartureTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 19)).
		SaveX(ctx)
	count++
	client.CallingPoint.
		Create().
		SetPlatform(PMS_1).
		SetArrivalTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 22)).
		SetDepartureTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 24)).
		SaveX(ctx)
	count++
	client.CallingPoint.
		Create().
		SetPlatform(FTN_1).
		SetArrivalTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 27)).
		SetDepartureTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 28)).
		SaveX(ctx)
	count++
	client.CallingPoint.
		Create().
		SetPlatform(HLS_1).
		SetArrivalTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 32)).
		SetDepartureTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 32)).
		SaveX(ctx)
	count++
	client.CallingPoint.
		Create().
		SetPlatform(BDH_1).
		SetArrivalTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 37)).
		SetDepartureTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 37)).
		SaveX(ctx)
	count++
	client.CallingPoint.
		Create().
		SetPlatform(HAV_1).
		SetArrivalTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 39)).
		SetDepartureTime(days[0].Date.Add(time.Hour * 8).Add(time.Minute * 40)).
		SaveX(ctx)
	count++

	log.Println("CallingPoints created:", count)
}
