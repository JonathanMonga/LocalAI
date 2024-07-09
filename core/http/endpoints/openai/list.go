package openai

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mudler/LocalAI/core/schema"
	"github.com/mudler/LocalAI/core/services"
)

// ListModelsEndpoint is the OpenAI Models API endpoint https://platform.openai.com/docs/api-reference/models
// @Summary List and describe the various models available in the API.
// @Success 200 {object} schema.ModelsDataResponse "Response"
// @Router /v1/models [get]
func ListModelsEndpoint(lms *services.ListModelsService) func(ctx *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// If blank, no filter is applied.
		filter := c.Query("filter")

		// By default, exclude any loose files that are already referenced by a configuration file.
		excludeConfigured := c.QueryBool("excludeConfigured", true)

		dataModels, err := lms.ListModels(filter, excludeConfigured)
		if err != nil {
			return err
		}
		return c.JSON(schema.ModelsDataResponse{
			Object: "list",
			Data:   dataModels,
		})
	}
}
