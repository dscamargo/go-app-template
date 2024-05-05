package controllers

import (
	"context"
	"github.com/dscamargo/go_app_template/internal/services"
	"github.com/dscamargo/go_app_template/pkg/web"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"time"
)

type GetAllExampleController struct {
	svc *services.ExampleService
}

func NewGetAllExampleController(svc *services.ExampleService) *GetAllExampleController {
	return &GetAllExampleController{svc: svc}
}

func (c *GetAllExampleController) Execute(fctx *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	items, err := c.svc.GetAll(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "GetAllExampleController.Execute", "error", err.Error())
		return web.NewInternalServerError(fctx)
	}

	return web.NewOKResponse(fctx, items)
}
