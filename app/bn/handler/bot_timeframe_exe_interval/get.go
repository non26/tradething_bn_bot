package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *botTimeframeExeIntervalHandler) GetHandler(c echo.Context) error {
	response, err := h.process.GetBotTimeframeExeInterval().Get(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
