package handlers

import (
	"{{.ProjectName}}/internal/services"
)

type Handlers struct {
	Example *ExampleHandler
}

func New(services *services.Services) *Handlers {
	return &Handlers{
		Example: NewExampleHandler(services.Example),
	}
}