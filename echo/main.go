package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main (){
	e := echo.New()
	e.Debug = true
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:"method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n"}))
	e.GET("/", func(c echo.Context)error{
		return c.String(http.StatusOK,"Hello,World")
	})
	e.Logger.Fatal(e.Start(":1818"))
}