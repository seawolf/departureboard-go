package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities"
	"bitbucket.org/sea_wolf/departure_board-go/v2/entities/service"
	"github.com/gin-gonic/gin"
)

func handleDepartureBoardService(c *gin.Context) {
	serviceIdStr, isFound := c.Params.Get("id")
	if !isFound {
		c.String(http.StatusNotFound, "Service ID not given")

		return
	}

	serviceId, err := strconv.Atoi(serviceIdStr)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())

		return
	}

	data, err := departureBoardService(CTX, Client, serviceId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())

		return
	}

	c.HTML(http.StatusOK, "service.tmpl", data)
}

func departureBoardService(ctx context.Context, client *entities.Client, serviceId int) (gin.H, error) {
	service, err := departureBoardServiceData(ctx, client, serviceId)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"Service": service,
	}, nil
}

func departureBoardServiceData(ctx context.Context, client *entities.Client, id int) (s *entities.Service, err error) {
	s, err = client.Service.
		Query().
		WithToc().
		WithDay().
		WithCallingPoints(func(cpq *entities.CallingPointQuery) {
			cpq.WithPlatform(func(pq *entities.PlatformQuery) {
				pq.WithStation()
			})
		}).
		Where(service.ID(id)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying Services: %w", err)
	}

	log.Println("Service returned:", s)
	return
}
