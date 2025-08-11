package handler

import (
	"net/http"
	"tradethingbot/app/bn/handler/req"

	"github.com/labstack/echo/v4"
)

func (h *botTimeframeExeIntervalHandler) HandlerSetBotTimeframeExeInterval(c echo.Context) error {
	request := &req.BotTimeframeExeIntervalHandlerRequest{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Transform(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.process.SetBotTimeframeExeInterval(c.Request().Context(), request.ToBotServiceRequest())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
