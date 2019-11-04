package api

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type LoginPayload struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserInfo struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	NotifyCount int    `json:"notifyCount"`
	UnreadCount int    `json:"unreadCount"`
	Avatar      string `json:"avatar"`
}

var loggedUser string

func UserLogin(c echo.Context) error {
	var loginPayload LoginPayload
	if err := c.Bind(&loginPayload); err != nil {
		return err
	}

	fmt.Printf("%+v\n", loginPayload)

	cookie := new(http.Cookie)
	cookie.Name = "appen_auth_session"
	cookie.Value = "OgCzQKa9VWZwSM7QhW4ZG4o2QjC78mlmguL="

	c.SetCookie(cookie)

	if loginPayload.UserName == "worker" && loginPayload.Password == "worker" {
		loggedUser = loginPayload.UserName

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  200,
			"message": "",
			"data": map[string]interface{}{
				"role": "worker",
				"type": "account",
			},
		})
	}

	if loginPayload.UserName == "manager" && loginPayload.Password == "manager" {
		loggedUser = loginPayload.UserName

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  200,
			"message": "",
			"data": map[string]interface{}{
				"role": "pm",
				"type": "account",
			},
		})
	}

	if loginPayload.UserName == "admin" && loginPayload.Password == "admin" {
		loggedUser = loginPayload.UserName

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  200,
			"message": "",
			"data": map[string]interface{}{
				"role": "admin",
				"type": "account",
			},
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  500,
		"message": "invalid username or password",
		"data": map[string]interface{}{
			"type": "account",
			"role": "guest",
		},
	})
}

func UserLogout(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func GetCurrentUser(c echo.Context) error {
	userInfo := UserInfo{
		Name:        loggedUser,
		Email:       loggedUser + "@appen.com",
		NotifyCount: 0,
		UnreadCount: 0,
		Avatar:      "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data":   userInfo,
	})
}

var wechatLoginStatus = map[string]bool{}

func GetWechatQRCode(c echo.Context) error {
	wechatUUID := c.QueryParam("uuid")

	wechatLoginStatus[wechatUUID] = false
	timer := time.NewTimer(5 * time.Second)

	go func(key string) {
		<-timer.C
		wechatLoginStatus[key] = true
	}(wechatUUID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]string{
			"url": gofakeit.URL(),
		},
	})
}

func GetWechatLoginStatus(c echo.Context) error {
	wechatUUID := c.QueryParam("uuid")

	loginStatus := wechatLoginStatus[wechatUUID]

	//	TODO if it's not first login, the data must include user name. and change login status in redux state of UI
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"uuid":       wechatUUID,
			"status":     loginStatus,
			"alias":      "Wechat Alias",
			"firstLogin": true,
		},
	})
}
