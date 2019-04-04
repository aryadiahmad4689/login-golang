package controllers

import (
	"fmt"
	"login-golang/db"
	"login-golang/models"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// index Login
func IndexLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login", map[string]interface{}{})
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginSave(c echo.Context) error {
	db := db.DbManager()
	user := []models.User{}
	password := c.FormValue("password")
	usernames := c.FormValue("username")
	db.Table("users").Select("*").Scan(&user)
	for _, v := range user {

		username := fmt.Sprintf("%s", v.Username)
		pass := fmt.Sprintf("%s", v.Password)

		match := CheckPasswordHash(password, pass)
		if username == usernames && match == true {
			return c.JSON(http.StatusOK, user)
		} else {
			return c.JSON(http.StatusOK, "Password Dan Email Anda Salah")
		}

	}
	return nil
}
