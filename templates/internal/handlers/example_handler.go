package handlers

import (
	"net/http"

	"{{.ProjectName}}/internal/services"
	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	exampleService *services.ExampleService
}

func NewExampleHandler(exampleService *services.ExampleService) *ExampleHandler {
	return &ExampleHandler{
		exampleService: exampleService,
	}
}

func (h *ExampleHandler) GetExample(c *gin.Context) {
	data := h.exampleService.GetExample()
	c.JSON(http.StatusOK, data)
}