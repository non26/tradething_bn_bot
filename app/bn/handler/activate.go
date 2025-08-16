package handler

import (
	"net/http"
	"tradethingbot/app/bn/handler/req"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/process"

	"github.com/labstack/echo/v4"
)

type activateHandler struct {
	process process.IBotService
}

func NewActivateHandler(process process.IBotService) *activateHandler {
	return &activateHandler{process: process}
}

func (h *activateHandler) HandleActivate(c echo.Context) error {
	request := req.ActivationRequestList{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	responses := make([]res.ActivationResponse, 0)
	for _, activation := range request.ActivationRequest {
		response := h.process.ActivateBot(c.Request().Context(), activation.ToDomain())
		responses = append(responses, *response)
	}

	return c.JSON(http.StatusOK, responses)
}
