package api

import (
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Job struct {
	JobId                 string    `json:"jobId"`
	JobType               string    `json:"jobType"`
	Name                  string    `json:"name"`
	Project               Project   `json:"project"`
	PrecedingJobId        string    `json:precedingJobId`
	DatasetInputLocation  string    `json:"datasetInputLocation"`
	DatasetOutputLocation string    `json:"datasetOutputLocation"`
	Creator               string    `json:"creator"`
	Category              string    `json:"category"`
	WorkerNumPerTask      int       `json:"workerNumPerTask"`
	RecordNumPerTask      int       `json:"recordNumPerTask"`
	CreatedAt             time.Time `json:"createdAt"`
	Status                string    `json:"status"`
}

func init() {
	for i := 0; i < 5; i++ {
		jobs = append(jobs, Job{
			JobId: gofakeit.UUID(),
			Name:  gofakeit.Address().Address,
			Project: Project{
				Name:      "Project: " + gofakeit.Name(),
				ProjectId: gofakeit.UUID(),
			},
			Creator:   gofakeit.Email(),
			Category:  "labeling_job",
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
		JobId: gofakeit.UUID(),
		Name:  gofakeit.Address().Address,
		Project: Project{
			Name:      "Project: " + gofakeit.Name(),
			ProjectId: gofakeit.UUID(),
		},
		Creator:   "Yingying",
		CreatedAt: time.Time{},
		Status:    "draft",
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   job,
	})
}

/*
dataset_input_location: "test"
dataset_output_location: "test"
job_name: "test"
job_type: "labeling_job"
labeling_category: "TRANSCRIPTION"
preceding_job_id: "2"
project_id: "f93fe66e-5ae4-4b59-89be-15f3fa494780"
record_num_per_task: "10"
template_id: "80dfe301-db41-40b9-94d5-d2ec38fe95db"
worker_num_per_task: "10"
*/
type JobPayload struct {
	JobName               string `json:"job_name"`
	JobType               string `json:"job_type"`
	ProjectId             string `json:"project_id"`
	PrecedingJobId        string `json:"preceding_job_id"`
	DatasetInputLocation  string `json:dataset_input_location`
	DatasetOutputLocation string `json:"dataset_output_location"`
	TemplateCategory      string `json:"labeling_category"`
	TemplateId            string `json:"template_id"`
	RecordNumPerTask      int    `json:"record_num_per_task"`
	WorkerNumPerTask      int    `json:"worker_num_per_task"`
}

func CreateJob(c echo.Context) error {
	var jobPayload JobPayload
	if err := c.Bind(&jobPayload); err != nil {
		return err
	}

	job := Job{
		JobId:   gofakeit.UUID(),
		JobType: jobPayload.JobType,
		Name:    jobPayload.JobName,
		Project: Project{
			Name:      "Project: " + gofakeit.Name(),
			ProjectId: jobPayload.ProjectId,
		},
		Creator:               "Yingjie",
		CreatedAt:             time.Time{},
		Status:                "draft",
		Category:              jobPayload.TemplateCategory,
		PrecedingJobId:        jobPayload.PrecedingJobId,
		DatasetInputLocation:  jobPayload.DatasetInputLocation,
		DatasetOutputLocation: jobPayload.DatasetOutputLocation,
		WorkerNumPerTask:      jobPayload.WorkerNumPerTask,
		RecordNumPerTask:      jobPayload.RecordNumPerTask,
	}

	jobs = append(jobs, job)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   job,
	})
}

func UpdateJob(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"jobId":   "86c93a59-91f0-41ae-9c8e-1ab5016ba837",
			"name":    "New Job Hello",
			"creator": "yj",
			"status":  "draft",
		},
	})
}
