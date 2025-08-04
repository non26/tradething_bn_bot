package handler

import (
	"net/http"
	"tradethingbot/app/bn/handler/req"
	"tradethingbot/app/bn/process"

	"github.com/labstack/echo/v4"
)

type invalidateBotHandler struct {
	process process.IBotService
}

func NewInvalidateBotHandler(
	process process.IBotService,
) *invalidateBotHandler {
	return &invalidateBotHandler{
		process: process,
	}
}

func (i *invalidateBotHandler) Handle(c echo.Context) error {
	request := &req.InvalidateBotHandlerRequest{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := i.process.InvalidateBot(c.Request().Context(), request.ToServiceModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
