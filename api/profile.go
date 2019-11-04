package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserProfile struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Mobile         string `json:"mobile"`
	IdentityCardNo string `json:"identityCardNo"`
}

var userProfile UserProfile

func init() {
	userProfile = UserProfile{
		Name:           "tyy",
		Email:          "555@gdygdyw.mmm",
		Mobile:         "15444444444",
		IdentityCardNo: "7777777****66666",
	}
}

func GetUserProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   userProfile,
	})
}

func UpdateUserProfile(c echo.Context) error {
	// TODO get user profile by id or email...
	var newUserProfile UserProfile
	if err := c.Bind(&newUserProfile); err != nil {
		return nil
	}

	userProfile = newUserProfile

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   newUserProfile,
	})
}
