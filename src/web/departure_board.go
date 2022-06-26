package web

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities"
	"github.com/gin-gonic/gin"
)

func handleDepartureBoardData(c *gin.Context) {
	data, err := departureBoardData(CTX, Client)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())

		return
	}

	c.HTML(http.StatusOK, "data.tmpl", data)
}

func departureBoardData(ctx context.Context, client *entities.Client) (gin.H, error) {
	tocs, err := departureBoardDataTOCs(ctx, client)
	if err != nil {
		return nil, err
	}

	days, err := departureBoardDataDays(ctx, client)
	if err != nil {
		return nil, err
	}

	stations, err := departureBoardDataStations(ctx, client)
	if err != nil {
		return nil, err
	}

	platforms, err := departureBoardDataPlatforms(ctx, client)
	if err != nil {
		return nil, err
	}

	callingPoints, err := departureBoardDataCallingPoints(ctx, client)
	if err != nil {
		return nil, err
	}

	services, err := departureBoardDataServices(ctx, client)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"TOCs":          tocs,
		"Days":          days,
		"Stations":      stations,
		"Platforms":     platforms,
		"CallingPoints": callingPoints,
		"Services":      services,
	}, nil
}

func departureBoardDataTOCs(ctx context.Context, client *entities.Client) (tocs []*entities.TOC, err error) {
	tocs, err = client.TOC.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying TOCs: %w", err)
	}

	log.Println("TOCs returned:", tocs)
	return
}

func departureBoardDataStations(ctx context.Context, client *entities.Client) (stations []*entities.Station, err error) {
	stations, err = client.Station.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying Stations: %w", err)
	}

	log.Println("Stations returned:", stations)
	return
}

func departureBoardDataPlatforms(ctx context.Context, client *entities.Client) (platforms []*entities.Platform, err error) {
	platforms, err = client.Platform.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying Platforms: %w", err)
	}

	log.Println("Platforms returned:", platforms)
	return
}

func departureBoardDataDays(ctx context.Context, client *entities.Client) (days []*entities.Day, err error) {
	days, err = client.Day.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying Days: %w", err)
	}

	log.Println("Days returned:", days)
	return
}

func departureBoardDataCallingPoints(ctx context.Context, client *entities.Client) (callingPoints []*entities.CallingPoint, err error) {
	callingPoints, err = client.CallingPoint.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying CallingPoints: %w", err)
	}

	log.Println("CallingPoints returned:", callingPoints)
	return
}

func departureBoardDataServices(ctx context.Context, client *entities.Client) (services []*entities.Service, err error) {
	services, err = client.Service.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying Services: %w", err)
	}

	log.Println("Services returned:", services)
	return
}