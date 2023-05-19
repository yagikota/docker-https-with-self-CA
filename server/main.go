package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yagikota/docker-https-with-self-CA/server/config"
)

type Response struct {
	Message string `json:"message"`
}

func greet(c echo.Context) error {
	message := fmt.Sprintf("Hello, World!, Current Time:%s", time.Now().Format("2006/01/02 15:04:05"))
	response := Response{Message: message}
	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", greet)

	cfg := config.LoadConfig()
	if err := http.ListenAndServeTLS(":443", cfg.TLS.CertFile, cfg.TLS.KeyFile, e); err != nil {
		log.Fatal(err)
	}
	log.Println("start HTTPS server 🐶")
}
