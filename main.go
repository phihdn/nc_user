package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phihdn/nc_user/config"
	mw "github.com/phihdn/nc_user/middleware"
	"github.com/phihdn/nc_user/route"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(mw.SimpleLogger())
	route.All(e)

	log.Println(e.Start(":9090"))
}
