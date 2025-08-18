package handler

import (
	"net/http"
	"tradethingbot/app/bn/handler/req"
	"tradethingbot/app/bn/handler/res"
	"tradethingbot/app/bn/process"

	"github.com/labstack/echo/v4"
)

type tvActivationHandler struct {
	process process.IBotService
}

func NewTVActivationHandler(process process.IBotService) *tvActivationHandler {
	return &tvActivationHandler{process: process}
}

func (h *tvActivationHandler) HandleTVActivation(c echo.Context) error {
	request := req.TVActivationRequest{}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	response := []res.ActivationResponse{}
	// deactivate bot
	for _, deactivate := range request.Deactivate {
		response1 := h.process.DeactivateBot(c.Request().Context(), deactivate.ToDomain())
		response = append(response, *response1)
	}

	// activate bot
	for _, activate := range request.Activate {
		response1 := h.process.ActivateBot(c.Request().Context(), activate.ToDomain())
		response = append(response, *response1)
	}

	return c.JSON(http.StatusOK, response)
}
