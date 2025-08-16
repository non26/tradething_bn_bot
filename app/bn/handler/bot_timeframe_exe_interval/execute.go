package handler

import (
	"net/http"
	"tradethingbot/app/bn/handler/req"

	"github.com/labstack/echo/v4"
)

func (h *botTimeframeExeIntervalHandler) ExecuteHandler(c echo.Context) error {
	request := &req.BotTimeframeExeIntervalHandlerRequest{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	h.process.DelayBot(ctx, h.delayTime)

	response, err := h.process.GetBotTimeframeExeInterval().Execute(ctx, request.ToBotServiceRequest())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
