package routers

import (
	"login-golang/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// exported Init
func Init() *echo.Echo {

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
	// ADMIN
	// admin := e.Group("/admin", middleware.BasicAuth(controllers.Login))
	//CORS

	// ROUTING
	e.GET("/register", controllers.IndexRegister)
	e.POST("/save", controllers.RegisterSave)
	e.GET("/login", controllers.IndexLogin)
	e.POST("/match", controllers.LoginSave)
	return e
}
