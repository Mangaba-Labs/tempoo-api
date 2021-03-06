package handler

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/weather"
	"github.com/gofiber/fiber/v2"
)

// WeatherHandler handler for /weather endpoints
type WeatherHandler interface {
	GetWeather(c *fiber.Ctx) error
}

// ServiceHandler struct
type ServiceHandler struct {
	service weather.WeatherService
}

// NewWeatherHandler initializer
func NewWeatherHandler(s weather.WeatherService) WeatherHandler {
	return &ServiceHandler{
		service: s,
	}
}

// GetWeather handler for GET /weather/current
func (h *ServiceHandler) GetWeather(c *fiber.Ctx) error {
	var userCoord = &weather.Request{}
	var service = weather.NewWeatherService()

	lat := c.Query("lat")
	lon := c.Query("lon")

	if lat == "" || lon == "" {
		return c.JSON(fiber.Map{"status": "error", "message": "malformed get-weather request", "data": nil})
	}

	userCoord.Latitude = lat
	userCoord.Longitude = lon

	currentWeather, err := service.GetCurrentWeather(userCoord)

	if err != nil {
		if err.Error() == "400" {
			c.Context().Response.SetStatusCode(400)
			return c.JSON(fiber.Map{"status": "error", "message": "malformed get-weather request", "data": nil})
		}
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": nil, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "data": currentWeather})
}
