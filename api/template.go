package api

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const (
	CONTENT_RELEVANCE = "CONTENT_RELEVANCE"
	SIDE_BY_SIDE      = "SIDE_BY_SIDE"
	TRANSCRIPTION     = "TRANSCRIPTION"
	TRANSLATION       = "TRANSLATION"
	TEXT_ANNOTATION   = "TEXT_ANNOTATION"
	IMAGE_ANNOTATION  = "IMAGE_ANNOTATION"
	LIDAR_ANNOTATION  = "LIDAR_ANNOTATION"
)

type Template struct {
	Id           string `json:"id"`
	Version      int    `json:"version"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Content      string `json:"content"`
	Instruction  string `json:"instruction"`
	Css          string `json:"css"`
	Options      string `json:"options"`
	TemplateType string `json:"type"`
	JobId        string `json:"jobId"`
}

type SimpleTemplate struct {
	Id           string `json:"id"`
	Version      int    `json:"version"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	TemplateType string `json:"type"`
}

func GetPublicTemplateTypes(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": []string{
			CONTENT_RELEVANCE,
			SIDE_BY_SIDE,
			TRANSCRIPTION,
			TRANSLATION,
			TEXT_ANNOTATION,
			IMAGE_ANNOTATION,
			LIDAR_ANNOTATION,
		},
	})
}

func GetPublicTemplatesByType(c echo.Context) error {
	// tmplType := c.Param("type")
	// var tmpls []SimpleTemplate

	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	tmplFolder := dir + "/static"

	files, err := readTemplateFiles(tmplFolder)

	if err != nil {
		return err
	}

	tmpls := parseTemplateJsonFiles(files)

	templateId := c.QueryParam("templateId")

	if templateId != "" {
		for i := 0; i < len(tmpls); i++ {
			if tmpls[i].Id == templateId {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"status": 200,
					"data":   tmpls[i],
				})
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   tmpls,
	})
}

var savedTemplates = map[string]Template{}

func SubmitTemplateByJobId(c echo.Context) error {
	var tmpl Template
	jobId := c.QueryParam("jobId")

	if err := c.Bind(&tmpl); err != nil {
		return err
	}

	fmt.Printf("************************* %s\n", jobId)
	fmt.Printf("%+v\n", tmpl)
	fmt.Println("-------------------------------")
	savedTemplates[jobId] = tmpl

	return c.NoContent(http.StatusOK)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": tmpl,
	})

}

func GetTemplateByJobId(c echo.Context) error {
	jobId := c.QueryParam("jobId")

	if jobId != "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 200,
			"data":   savedTemplates[jobId],
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   savedTemplates,
	})
}

func readTemplateFiles(folderPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".json" {
			return nil
		}

		files = append(files, path)

		return nil
	})

	if err != nil {
		return files, err
	}

	return files, nil
}

func parseTemplateJsonFiles(files []string) []Template {
	var templates []Template
	for i := 0; i < len(files); i++ {
		jsonBytes, err := ioutil.ReadFile(files[i])

		if err != nil {
			log.Debug(err)
		}

		var tmpl Template
		err = json.Unmarshal(jsonBytes, &tmpl)
		if err == nil {
			templates = append(templates, tmpl)
		} else {
			log.Debug(err)
		}
	}

	return templates
}
