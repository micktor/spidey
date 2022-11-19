package httpd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/micktor/spidey/internal"
	"github.com/micktor/spidey/internal/config"
	"github.com/micktor/spidey/internal/dto"
	"github.com/pkg/errors"
)

// NewSpideyHandler returns a configured PaymentMethodHandler handler.
func NewSpideyHandler(applicationConfig *applicationConfig) *SpideyHandler {
	return &SpideyHandler{
		spideyService: applicationConfig.spideyService,
	}
}

type SpideyHandler struct {
	config        *config.Config
	spideyService internal.SpideyService
}

func (t SpideyHandler) CreateSpidey(c *fiber.Ctx) error {
	request := new(dto.SpideyRequest)
	if err := c.BodyParser(request); err != nil {
		c.SendStatus(fiber.StatusServiceUnavailable)
		return nil
	}

	resp, err := t.spideyService.Test(*request)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return nil
	}

	switch errors.Cause(err) {
	case nil:
		c.JSON(resp)
	default:
		c.SendStatus(fiber.StatusInternalServerError)
	}
	return nil
}
