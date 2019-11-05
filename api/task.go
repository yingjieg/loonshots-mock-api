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
	JobId       string    `json:"jobId"`
	JobName     string    `json:"jobName"`
	JobCategory string    `json:"jobCategory"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Status      string    `json:"status"`
}

type TaskRecord struct {
	Id     string                 `json:"id"`
	Fields map[string]interface{} `json:"fields"`
}

func GetTasks(c echo.Context) error {
	var tasks []Task
	for i := 0; i < len(jobs); i++ {
		tasks = append(tasks, Task{
			Id:          gofakeit.UUID(),
			ProjectName: gofakeit.Name(),
			JobId:       jobs[i].JobId,
			JobName:     jobs[i].Name,
			JobCategory: jobs[i].Category,
			StartTime:   time.Time{},
			EndTime:     time.Time{},
			Status:      gofakeit.RandString([]string{"TODO", "DONE"}),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"totalElement": len(tasks),
			"content":      tasks,
		},
	})
}

func GetTaskRecords(c echo.Context) error {
	taskId := c.QueryParam("taskId")

	if taskId == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "bad request, taskId not found",
		})
	}

	var taskRecords []TaskRecord
	taskRecords = append(taskRecords, TaskRecord{
		Id: "0241008287272164729465721528295504357920",
		Fields: map[string]interface{}{
			"name": "jim",
			"test": "hello",
		},
	}, TaskRecord{
		Id: "0241008287272164729465721528295504357921",
		Fields: map[string]interface{}{
			"name": "james",
			"test": "world",
		},
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   taskRecords,
	})

}
