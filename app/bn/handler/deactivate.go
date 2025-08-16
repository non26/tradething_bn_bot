package handler

import (
	"net/http"
	"tradethingbot/app/bn/handler/req"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/process"

	"github.com/labstack/echo/v4"
)

type deactivateHandler struct {
	process process.IBotService
}

func NewDeactivateHandler(process process.IBotService) *deactivateHandler {
	return &deactivateHandler{process: process}
}

func (h *deactivateHandler) HandleDeactivate(c echo.Context) error {
	request := req.ActivationRequestList{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	responses := make([]res.ActivationResponse, 0)
	for _, deactivation := range request.ActivationRequest {
		response := h.process.DeactivateBot(c.Request().Context(), deactivation.ToDomain())
		responses = append(responses, *response)
	}

	return c.JSON(http.StatusOK, responses)
}
