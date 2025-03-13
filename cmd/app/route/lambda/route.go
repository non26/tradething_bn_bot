package route

import (
	"tradethingbot/config"

	"github.com/labstack/echo/v4"
)

func RouteLambda(
	app *echo.Echo, config *config.AppConfig,
) {
	UpdateAWSAppConfig(app, config)
}
