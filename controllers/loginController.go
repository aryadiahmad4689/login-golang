package controllers

import (
	"fmt"
	"login-golang/db"
	"login-golang/models"
	"net/http"

	hash "github.com/aryadiahmad4689/hash-bycript"

	"github.com/labstack/echo"
)

// index Login
func IndexLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login", map[string]interface{}{})
}

// func Login
func LoginSave(c echo.Context) error {
	db := db.DbManager()
	user := []models.User{}
	password := c.FormValue("password")
	usernames := c.FormValue("username")
	db.Table("users").Select("*").Scan(&user)
	for _, v := range user {
		username := fmt.Sprintf("%s", v.Username)
		pass := fmt.Sprintf("%s", v.Password)
		match := hash.CheckPasswordHash(password, pass)
		if username == usernames && match == true {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.JSON(http.StatusOK, "Password Dan Username Salah")
}
