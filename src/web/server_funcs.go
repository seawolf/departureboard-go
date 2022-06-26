package web

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"bitbucket.org/sea_wolf/departure_board-go/v2/entities"
	"github.com/gin-gonic/gin"
)

func funcs(engine *gin.Engine) {
	engine.SetFuncMap(template.FuncMap{
		"formatAsDate":      formatAsDate,
		"formatAsTime":      formatAsTime,
		"listCallingPoints": listCallingPoints,
	})
}

func formatAsDate(t time.Time) string {
	if t.IsZero() {
		return "-"
	}

	return fmt.Sprintf("%02d/%02d/%04d", t.Day(), t.Month(), t.Year())
}

func formatAsTime(t time.Time) string {
	if t.IsZero() {
		return "-"
	}

	return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
}

func listCallingPoints(callingPoints []*entities.CallingPoint) template.HTML {
	list := make([]string, len(callingPoints))

	for idx, callingPoint := range callingPoints {
		list[idx] = callingPoint.Edges.Platform.Edges.Station.Name
	}

	return template.HTML(strings.Join(list[:], "<br />"))
}
