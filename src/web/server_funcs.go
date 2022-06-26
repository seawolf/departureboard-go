package web

import (
	"fmt"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

func funcs(engine *gin.Engine) {
	engine.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
		"formatAsTime": formatAsTime,
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
