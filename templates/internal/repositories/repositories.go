package repositories

import (
	"gorm.io/gorm"
)

type Repositories struct {
	Example *ExampleRepository
}

func New(db *gorm.DB) *Repositories {
	return &Repositories{
		Example: NewExampleRepository(db),
	}
}