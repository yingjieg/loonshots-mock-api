package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"loonshots-mock-api/api"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/* --------------- Routes ----------------- */

	// user auth
	e.POST("/api/auth/v1/login", api.UserLogin)
	e.POST("/api/auth/v1/logout", api.UserLogout)
	e.GET("/api/auth/v1/currentUser", api.GetCurrentUser)

	e.GET("/api/auth/v1/wechat-qr", api.GetWechatQRCode)
	e.GET("/api/auth/v1/wechat-status", api.GetWechatLoginStatus)

	// user register
	e.POST("/api/auth/v1/register", api.UserRegister)

	// user profile
	e.GET("/api/user/v1/user-profile", api.GetUserProfile)
	e.PUT("/api/user/v1/user-profile", api.UpdateUserProfile)

	// project
	e.GET("/api/v1/projects", api.GetProjects)
	e.POST("/api/v1/projects", api.CreateProject)

	// job
	e.GET("/api/v1/jobs", api.GetJobs)
	e.GET("/api/v1/job", api.GetJob)
	e.POST("/api/v1/jobs", api.CreateJob)
	e.PUT("/api/v1/jobs", api.UpdateJob)

	// template
	e.GET("/api/v1/public_template_types", api.GetPublicTemplateTypes)
	e.GET("/api/v1/public_templates", api.GetPublicTemplatesByType)
	e.GET("/api/v1/templates", api.GetTemplateByJobId)
	e.POST("/api/v1/templates", api.SubmitTemplateByJobId)

	// dataset
	e.GET("/api/v1/dataset", api.GetDatasets)
	e.POST("/api/v1/dataset-upload", api.FileUpload)
	e.GET("/api/v1/dataset-download", api.FileDownload)

	// task
	e.GET("/api/v1/tasks", api.GetTasks)

	/* ------------------------------------------------- */

	// Start server
	e.Logger.Fatal(e.Start(":3002"))
}
