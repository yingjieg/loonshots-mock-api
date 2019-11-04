package api

import (
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	dst, err := os.Create(filepath.Join("files", file.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"createdAt": gofakeit.Date(),
			"creator":   gofakeit.Username(),
			"location":  filepath.Join("files", file.Filename),
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

	return c.Attachment(dataset.Location, "dataset.csv")
}
