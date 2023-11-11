package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikochan01/WeatherAPP-API/weather" // Ganti dengan path yang sesuai
)

func main() {
	e := echo.New()

	// Endpoint untuk mendapatkan informasi cuaca
	e.GET("/weather", func(c echo.Context) error {
		result, err := weather.GetWeather()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, result)
	})

	e.Start(":8080")
}
