package api

import (
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo/v4"
	"loonshots-mock-api/utils"
	"net/http"
	"time"
)

type Dataset struct {
	Id        string    `json:"id"`
	Location  string    `json:"location"`
	Creator   string    `json:"creator"`
	CreatedAt time.Time `json:"createdAt"`
	Size      int       `json:"size"`
}

var datasetList []Dataset

func GetDatasets(c echo.Context) error {

	files, _ := utils.ListFolder("files")

	datasetList = []Dataset{}
	for i := 0; i < len(files); i++ {
		datasetList = append(datasetList, Dataset{
			Id:        gofakeit.UUID(),
			Location:  files[i],
			Creator:   gofakeit.Username(),
			CreatedAt: time.Time{},
			Size:      gofakeit.Number(100, 10000),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"totalElement": len(datasetList),
			"content":      datasetList,
		},
	})
}

func PreviewDataset(c echo.Context) error {
	var rows [][]string
	for i := 0; i < 5; i++ {
		rows = append(rows, []string{
			gofakeit.Username(),
			gofakeit.URL(),
			gofakeit.Company(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"headers": []string{"username", "url", "company"},
			"rows":    rows,
		},
	})

}
