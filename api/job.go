package api

import (
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Job struct {
	JobId           string    `json:"id"`
	ProjectId       string    `json:"projectId"`
	JobType         string    `json:"jobType"`
	JobName         string    `json:"jobName"`
	PreconditionJob string    `json:"preconditionJob"`
	SourceFile      string    `json:"sourceFile"`
	TargetTopic     string    `json:"targetTopic"`
	DataSelection   int       `json:"dataSelection"`
	WorkerNum       int       `json:"workerNum"`
	RecordNum       int       `json:"recordNum"`
	CreatedAt       time.Time `json:"createdAt"`
	Owner           string    `json:"owner"`
	Status          string    `json:"status"`
	Tag             string    `json:"tag"`
}

func init() {
	for i := 0; i < 1; i++ {
		jobs = append(jobs, Job{
			JobId:     gofakeit.UUID(),
			JobType:   "labeling",
			JobName:   gofakeit.Address().Address,
			CreatedAt: gofakeit.Date(),
			Status:    gofakeit.RandString([]string{"draft", "launched", "started", "stopped"}),
		})
	}
}

var jobs []Job

func GetJobs(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"totalElements": len(jobs),
			"content":       jobs,
		},
	})
}
func GetJob(c echo.Context) error {

	jobId := c.QueryParam("jobId")

	for i := 0; i < len(jobs); i++ {
		if jobs[i].JobId == jobId {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"status": 200,
				"data":   jobs[i],
			})
		}
	}

	job := Job{
		JobId:     gofakeit.UUID(),
		JobName:   gofakeit.Address().Address,
		Owner:     "Yingying",
		CreatedAt: time.Time{},
		Status:    "draft",
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   job,
	})
}

type JobPayload struct {
	JobName          string `json:"jobName"`
	JobType          string `json:"jobType"`
	ProjectId        string `json:"projectId"`
	PreconditionJob  string `json:"preconditionJob"`
	SourceFile       string `json:"sourceFile"`
	TargetTopic      string `json:"targetTopic"`
	TemplateCategory string `json:"labelingCategory"`
	TemplateId       string `json:"templateId"`
	dataSelection    int    `json:"dataSelection"`
	RecordNum        int    `json:"recordNum"`
	WorkerNum        int    `json:"workerNum"`
	Tag              string `json:"tag"`
}

func CreateJob(c echo.Context) error {
	var jobPayload JobPayload
	if err := c.Bind(&jobPayload); err != nil {
		return err
	}

	job := Job{
		JobId:           gofakeit.UUID(),
		ProjectId:       jobPayload.ProjectId,
		JobType:         jobPayload.JobType,
		JobName:         jobPayload.JobName,
		Owner:           "Yingjie",
		CreatedAt:       time.Time{},
		Status:          "draft",
		PreconditionJob: jobPayload.PreconditionJob,
		SourceFile:      jobPayload.SourceFile,
		TargetTopic:     jobPayload.TargetTopic,
		WorkerNum:       jobPayload.WorkerNum,
		RecordNum:       jobPayload.RecordNum,
	}

	jobs = append(jobs, job)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   job,
	})
}

func UpdateJob(c echo.Context) error {
	jobId := c.QueryParam("jobId")

	var job Job
	if err := c.Bind(&job); err != nil {
		return err
	}

	for i := 0; i < len(jobs); i++ {
		if jobs[i].JobId == jobId {
			jobs[i] = job
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"id":      "86c93a59-91f0-41ae-9c8e-1ab5016ba837",
			"jobName": "New Job Hello",
			"owner":   "yj",
			"status":  "draft",
		},
	})
}
