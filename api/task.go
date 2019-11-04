package api

import (
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Task struct {
	Id          string    `json:"id"`
	ProjectName string    `json:"projectName"`
	JobName     string    `json:"jobName"`
	JobCategory string    `json:"jobCategory"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Status      string    `json:"status"`
}

func init() {
	tasks = append(tasks, Task{
		Id:          gofakeit.UUID(),
		ProjectName: gofakeit.Name(),
		JobName:     gofakeit.StreetName(),
		JobCategory: gofakeit.CarModel(),
		StartTime:   gofakeit.Date(),
		EndTime:     gofakeit.Date(),
		Status:      "TODO",
	}, Task{
		Id:          gofakeit.UUID(),
		ProjectName: gofakeit.Name(),
		JobName:     gofakeit.StreetName(),
		JobCategory: gofakeit.CarModel(),
		StartTime:   gofakeit.Date(),
		EndTime:     gofakeit.Date(),
		Status:      "DONE",
	})
}

var tasks []Task

func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"totalElement": len(tasks),
			"content":      tasks,
		},
	})
}
