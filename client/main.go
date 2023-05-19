package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	Message string `json:"message"`
}

func greet(c echo.Context) error {
	url := "https://server:443"
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		Get(url)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	var response Response
	// stream処理できないのか？
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", greet)
	e.Logger.Fatal(e.Start(":8081"))
}
