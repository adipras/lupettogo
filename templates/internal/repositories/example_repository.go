package repositories

import (
	"{{.ProjectName}}/internal/models"
	"gorm.io/gorm"
)

type ExampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) *ExampleRepository {
	return &ExampleRepository{
		db: db,
	}
}

func (r *ExampleRepository) FindAll() ([]*models.Example, error) {
	var examples []*models.Example
	err := r.db.Find(&examples).Error
	return examples, err
}

func (r *ExampleRepository) FindByID(id uint) (*models.Example, error) {
	var example models.Example
	err := r.db.First(&example, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &example, nil
}

func (r *ExampleRepository) Create(example *models.Example) (*models.Example, error) {
	err := r.db.Create(example).Error
	return example, err
}

func (r *ExampleRepository) Update(example *models.Example) (*models.Example, error) {
	err := r.db.Save(example).Error
	return example, err
}

func (r *ExampleRepository) Delete(id uint) error {
	return r.db.Delete(&models.Example{}, id).Error
}