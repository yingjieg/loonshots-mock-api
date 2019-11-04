package api

import (
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type ProjectPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	WorkdayId   string `json:"workdayId"`
}

type Project struct {
	Name        string    `json:"name"`
	ProjectId   string    `json:"projectId"`
	Logo        string    `json:"logo"`
	Description string    `json:"description"`
	WorkdayId   string    `json:"workdayId"`
	Creator     string    `json:"creator"`
	CreatedAt   time.Time `json:"createdAt"`
}

func init() {
	gofakeit.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		projects = append(projects, Project{
			Name:        gofakeit.Company(),
			ProjectId:   gofakeit.UUID(),
			Description: gofakeit.Quote(),
			WorkdayId:   "99288123",
			Creator:     gofakeit.Email(),
			CreatedAt:   gofakeit.Date(),
		})
	}

}

var projects []Project

func GetProjects(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   projects,
	})
}

func CreateProject(c echo.Context) error {
	var projectPayload ProjectPayload
	if err := c.Bind(&projectPayload); err != nil {
		return err
	}

	project := Project{
		Name:        projectPayload.Name,
		ProjectId:   gofakeit.UUID(),
		Description: gofakeit.Quote(),
		WorkdayId:   projectPayload.WorkdayId,
		Creator:     loggedUser,
		CreatedAt:   time.Time{},
	}

	projects = append(projects, project)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   project,
	})
}
