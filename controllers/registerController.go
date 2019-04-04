package controllers

import (
	"login-golang/db"
	"login-golang/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"

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
	b := []byte(c.FormValue("password"))
	pass, _ := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	user.Nama = c.FormValue("nama")
	user.Username = c.FormValue("username")
	user.Password = string(pass)
	db.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
