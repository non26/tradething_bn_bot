package handler

import (
	"net/http"
	"tradethingbot/app/bn/handler/req"

	"github.com/labstack/echo/v4"
)

func (h *botTimeframeExeIntervalHandler) SetHandler(c echo.Context) error {
	request := &req.SetBotTimeframeExeIntervalHandlerRequest{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.process.GetBotTimeframeExeInterval().Set(c.Request().Context(), request.ToBotServiceRequest())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
