package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type RegisterPayload struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func UserRegister(c echo.Context) error {
	var registerPayload RegisterPayload
	if err := c.Bind(&registerPayload); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]string{
			"email": registerPayload.Email,
			"name":  registerPayload.Name,
		},
	})
}
