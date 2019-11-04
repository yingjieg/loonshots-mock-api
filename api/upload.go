package api

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func FileUpload(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if _, err := os.Stat("files"); os.IsNotExist(err) {
		os.Mkdir("files", 0775)
	}

	// Destination
	dst, err := os.Create("files/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"createdAt": gofakeit.Date(),
			"creator":   gofakeit.Username(),
			"location":  "files/" + file.Filename,
			"size":      file.Size,
		},
	})

}

func FileDownload(c echo.Context) error {
	datasetId := c.QueryParam("datasetId")

	var dataset Dataset
	for i := 0; i < len(datasetList); i++ {
		if datasetList[i].Id == datasetId {
			dataset = datasetList[i]
		}
	}

	fmt.Printf("*****************\n")
	fmt.Println(dataset.Location)

	return c.Attachment(dataset.Location, "dataset.csv")
}
