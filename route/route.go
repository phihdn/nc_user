package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phihdn/nc_user/config"
	"github.com/phihdn/nc_user/handler"
	"github.com/phihdn/nc_user/models"
)

func All(e *echo.Echo) {
	Private(e)
	Public(e)
}

func Private(e *echo.Echo) {
	g := e.Group("/api/v1/private")
	JWTConfig := middleware.JWTConfig{
		SigningKey: []byte(config.Config.JWTSecret.JWTKey),
		Claims:     &models.UserClaims{},
	}
	g.Use(middleware.JWTWithConfig(JWTConfig))
	g.PUT("/user", handler.UpdateUser)
}

func Public(e *echo.Echo) {
	g := e.Group("/api/v1/public")
	g.GET("/health", handler.HealthCheck)
	g.POST("/user/register", handler.RegisterUser)
	g.PATCH("/user/login", handler.LoginUser)

}
