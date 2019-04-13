package controllers

import (
	"login-golang/db"
	"login-golang/models"
	"net/http"

	hash "github.com/aryadiahmad4689/hash-bycript"

	"github.com/labstack/echo"
)

// func index REGISTER
func IndexRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "index", map[string]interface{}{})
}

// func REGISTER SAVE
func RegisterSave(c echo.Context) error {
	db := db.DbManager()
	user := new(models.User)
	user.Nama = c.FormValue("nama")
	user.Username = c.FormValue("username")
	user.Password, _ = hash.HashPassword(c.FormValue("password"), 12)
	db.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
