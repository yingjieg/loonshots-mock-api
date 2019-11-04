package api

import (
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo/v4"
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

var datasets []Dataset

func init() {
	datasets = append(datasets, Dataset{
		Id:        gofakeit.UUID(),
		Location:  "s3://data.appen.com/storage//object/buckets/xyz/abx.d",
		Creator:   gofakeit.Email(),
		CreatedAt: time.Time{},
		Size:      gofakeit.Number(100, 10000),
	}, Dataset{
		Id:        gofakeit.UUID(),
		Location:  "s3://data.appen.com/storage/object/buckets/xyz/abx.d",
		Creator:   gofakeit.Email(),
		CreatedAt: time.Time{},
		Size:      gofakeit.Number(100, 10000),
	})
}

func GetDatasets(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"totalItem": len(datasets),
			"content":   datasets,
		},
	})
}

func UploadDataset(c echo.Context) error {
	return nil
}

func DownloadDataset(c echo.Context) error {
	return nil
}

func PreviewDataset(c echo.Context) error {
	return nil
}
